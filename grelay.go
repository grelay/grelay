package grelay

import (
	"context"
	"sync"
	"time"
)

type Grelay interface {
	Exec(func() (interface{}, error)) (interface{}, error)
}

type grelayImpl struct {
	config                   GrelayConfig
	currentServiceThreshould int64
	state                    state

	mu sync.RWMutex
}

type callResponse struct {
	i   interface{}
	err error
}

func NewGrelay(c GrelayConfig) Grelay {
	g := &grelayImpl{
		config: c,
		state:  closed,
	}
	go g.monitoring()
	return g
}

func (g *grelayImpl) Exec(f func() (interface{}, error)) (interface{}, error) {
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

func (g *grelayImpl) makeCall(f func() (interface{}, error), c chan<- callResponse) {
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

func (g *grelayImpl) monitoring() {
	for range time.Tick(g.config.retryTimePeriod) {
		g.mu.RLock()
		if string(g.state) == string(closed) {
			g.mu.RUnlock()
			continue
		}
		g.mu.RUnlock()

		g.mu.Lock()
		g.state = halfOpen

		ctx, cancel := context.WithTimeout(context.Background(), g.config.serviceTimeout)

		select {
		case <-ctx.Done():

			g.state = open
			cancel()
			g.mu.Unlock()
			continue
		default:
			if err := g.config.service.Ping(); err != nil {
				g.state = open
				cancel()
				g.mu.Unlock()
				continue
			}
			g.state = closed
			g.currentServiceThreshould = 0
			cancel()
			g.mu.Unlock()
		}
	}
}
