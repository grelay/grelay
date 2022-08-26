package gr

import (
	"sync"
	"time"
)

type GrelayService interface {
	exec(func() (interface{}, error)) (interface{}, error)
}

type grelayServiceImpl struct {
	state                    state
	currentServiceThreshould int64
	config                   GrelayConfig

	mu sync.RWMutex
}

type callResponse struct {
	i   interface{}
	err error
}

func NewGrelayService(c GrelayConfig) GrelayService {
	g := &grelayServiceImpl{
		config: c,
		state:  closed,
	}
	go g.monitoring()
	return g
}

func (g *grelayServiceImpl) exec(f func() (interface{}, error)) (interface{}, error) {
	return getGrelayExec(g.config).exec(g, f)
}

func (g *grelayServiceImpl) monitoring() {
	for range time.Tick(g.config.retryTimePeriod) {
		g.monitoringState()
	}
}

func (g *grelayServiceImpl) monitoringState() {
	g.mu.RLock()
	if string(g.state) == string(closed) {
		if g.currentServiceThreshould > 0 {
			g.mu.RUnlock()

			checkerChannel := make(chan bool, 1)
			go g.checkService(checkerChannel)

			g.mu.RLock()
			t := time.NewTimer(g.config.serviceTimeout)
			g.mu.RUnlock()
			defer t.Stop()

			select {
			case <-t.C:
				return
			case ok := <-checkerChannel:
				if !ok {
					return
				}
				g.mu.Lock()
				defer g.mu.Unlock()
				g.currentServiceThreshould = 0
				return
			}
		}
		g.mu.RUnlock()
		return
	}
	g.mu.RUnlock()

	g.mu.Lock()
	g.state = halfOpen
	g.mu.Unlock()

	checkerChannel := make(chan bool, 1)
	go g.checkService(checkerChannel)

	g.mu.RLock()
	t := time.NewTimer(g.config.serviceTimeout)
	g.mu.RUnlock()
	defer t.Stop()

	select {
	case <-t.C:
		g.mu.Lock()
		g.state = open
		g.mu.Unlock()
		return
	case ok := <-checkerChannel:
		g.mu.Lock()
		defer g.mu.Unlock()
		if !ok {
			g.state = open
			return
		}
		g.state = closed
		g.currentServiceThreshould = 0
		return
	}
}

func (g *grelayServiceImpl) checkService(c chan<- bool) {
	defer close(c)
	g.mu.RLock()
	defer g.mu.RUnlock()
	if err := g.config.service.Ping(); err != nil {
		c <- false
		return
	}
	c <- true
}
