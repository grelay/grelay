package grelay

import (
	"context"
	"sync"
	"time"
)

type GrelayService interface {
	exec(func() (interface{}, error)) (interface{}, error)
}

type grelayServiceImpl struct {
	config                   GrelayConfig
	currentServiceThreshould int64
	state                    state

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
	ctx, cancel := context.WithTimeout(context.Background(), g.config.serviceTimeout)
	defer cancel()

	callDone := make(chan callResponse, 1)
	go g.makeCall(f, callDone)

	select {
	case <-ctx.Done():
		g.mu.Lock()
		g.currentServiceThreshould++
		g.mu.Unlock()

		g.mu.RLock()
		if g.currentServiceThreshould >= g.config.serviceThreshould {
			g.mu.RUnlock()

			g.mu.Lock()
			g.state = open
			g.mu.Unlock()

			return nil, ErrGrelayServiceTimedout
		}
		g.mu.RUnlock()

		return nil, ErrGrelayServiceTimedout
	case r := <-callDone:
		return r.i, r.err
	}
}

func (g *grelayServiceImpl) makeCall(f func() (interface{}, error), c chan<- callResponse) {
	g.mu.RLock()
	if string(g.state) == string(open) || string(g.state) == string(halfOpen) {
		g.mu.RUnlock()
		c <- callResponse{nil, ErrGrelayStateOpened}
		return
	}
	if g.currentServiceThreshould >= g.config.serviceThreshould {
		g.mu.RUnlock()

		g.mu.Lock()
		g.state = open
		g.mu.Unlock()

		c <- callResponse{nil, ErrGrelayStateOpened}
		return
	}
	g.mu.RUnlock()

	i, err := f()
	c <- callResponse{i, err}
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
			ctx, cancel := context.WithTimeout(context.Background(), g.config.serviceTimeout)

			select {
			case <-ctx.Done():
				cancel()
				return
			case ok := <-checkerChannel:
				if !ok {
					cancel()
					return
				}
				g.mu.Lock()
				g.currentServiceThreshould = 0
				g.mu.Unlock()
				cancel()
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

	ctx, cancel := context.WithTimeout(context.Background(), g.config.serviceTimeout)

	select {
	case <-ctx.Done():
		g.mu.Lock()
		g.state = open
		cancel()
		g.mu.Unlock()
		return
	case ok := <-checkerChannel:
		g.mu.Lock()
		defer g.mu.Unlock()
		if !ok {
			g.state = open
			cancel()
			return
		}
		g.state = closed
		g.currentServiceThreshould = 0
		cancel()
		return
	}
}

func (g *grelayServiceImpl) checkService(c chan<- bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if err := g.config.service.Ping(); err != nil {
		c <- false
		return
	}
	c <- true
}
