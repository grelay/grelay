package grelay

import (
	"context"
	"sync"
	"time"

	"github.com/grelay/grelay/internal/states"
)

type Service struct {
	state                    string
	currentServiceThreshould int64
	config                   Configuration
	service                  Pingable

	mu *sync.RWMutex
}

type callResponse struct {
	i   interface{}
	err error
}

func NewGrelayService(cfg Configuration, service Pingable) *Service {
	g := &Service{
		config:  cfg,
		state:   states.Closed,
		service: service,

		mu: &sync.RWMutex{},
	}
	go g.monitoring()
	return g
}

func (g *Service) exec(f func() (interface{}, error)) (interface{}, error) {
	return getGrelayExec(g.config).exec(g, f)
}

func (g *Service) monitoring() {
	for range time.Tick(g.config.RetryPeriod) {
		g.monitoringState()
	}
}

func (g *Service) monitoringState() {
	g.mu.RLock()
	ctx, cancel := context.WithTimeout(context.Background(), g.config.Timeout)
	g.mu.RUnlock()
	defer cancel()

	g.mu.RLock()
	if g.state == states.Closed {
		if g.currentServiceThreshould > 0 {
			g.mu.RUnlock()

			checkerChannel := make(chan bool, 1)
			go g.checkService(checkerChannel)

			select {
			case <-ctx.Done():
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
	g.state = states.HalfOpen
	g.mu.Unlock()

	checkerChannel := make(chan bool, 1)
	go g.checkService(checkerChannel)

	select {
	case <-ctx.Done():
		g.mu.Lock()
		g.state = states.Open
		g.mu.Unlock()
		return
	case ok := <-checkerChannel:
		g.mu.Lock()
		defer g.mu.Unlock()
		if !ok {
			g.state = states.Open
			return
		}
		g.state = states.Closed
		g.currentServiceThreshould = 0
		return
	}
}

func (g *Service) checkService(c chan<- bool) {
	defer close(c)
	g.mu.RLock()
	defer g.mu.RUnlock()
	if err := g.service.Ping(); err != nil {
		c <- false
		return
	}
	c <- true
}
