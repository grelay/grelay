package grelay

import (
	"time"

	"github.com/grelay/grelay/internal/errs"
	"github.com/grelay/grelay/internal/states"
)

type grelayExecWithGo struct{}

func (e grelayExecWithGo) exec(service *Service, f func() (interface{}, error)) (interface{}, error) {
	gs := service
	callDone := make(chan callResponse, 1)
	go e.makeCall(gs, f, callDone)

	gs.mu.RLock()
	t := time.NewTimer(gs.config.Timeout)
	gs.mu.RUnlock()
	defer t.Stop()

	select {
	case <-t.C:
		gs.mu.Lock()
		gs.currentServiceThreshould++
		gs.mu.Unlock()

		gs.mu.RLock()
		if gs.currentServiceThreshould >= gs.config.Threshould {
			gs.mu.RUnlock()

			gs.mu.Lock()
			gs.state = states.Open
			gs.mu.Unlock()

			return nil, errs.ErrGrelayServiceTimedout
		}
		gs.mu.RUnlock()

		return nil, errs.ErrGrelayServiceTimedout
	case r := <-callDone:
		return r.i, r.err
	}
}

func (g grelayExecWithGo) makeCall(service *Service, f func() (interface{}, error), c chan<- callResponse) {
	gs := service
	defer close(c)
	gs.mu.RLock()
	if string(gs.state) == string(states.Open) || string(gs.state) == string(states.HalfOpen) {
		gs.mu.RUnlock()
		c <- callResponse{nil, errs.ErrGrelayStateOpened}
		return
	}
	if gs.currentServiceThreshould >= gs.config.Threshould {
		gs.mu.RUnlock()

		gs.mu.Lock()
		gs.state = states.Open
		gs.mu.Unlock()

		c <- callResponse{nil, errs.ErrGrelayStateOpened}
		return
	}
	gs.mu.RUnlock()

	i, err := f()
	c <- callResponse{i, err}
}
