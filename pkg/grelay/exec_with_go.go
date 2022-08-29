package grelay

import (
	"time"

	"github.com/grelay/grelay/internal/errs"
	"github.com/grelay/grelay/internal/states"
)

type grelayExecWithGo struct{}

func (e grelayExecWithGo) exec(service *Service, f func() (interface{}, error)) (interface{}, error) {
	service.mu.RLock()
	timeout := time.NewTimer(service.config.Timeout)
	service.mu.RUnlock()
	defer timeout.Stop()

	callDone := make(chan callResponse, 1)
	go e.makeCall(service, f, callDone)

	select {
	case <-timeout.C:
		return nil, e.processTimeout(service)
	case r := <-callDone:
		if !timeout.Stop() {
			<-timeout.C
			return nil, e.processTimeout(service)
		}
		return r.i, r.err
	}
}

func (g grelayExecWithGo) makeCall(service *Service, f func() (interface{}, error), c chan<- callResponse) {
	defer close(c)
	service.mu.RLock()
	if service.state == states.Open || service.state == states.HalfOpen {
		service.mu.RUnlock()
		c <- callResponse{nil, errs.ErrGrelayStateOpened}
		return
	}
	if service.currentServiceThreshould >= service.config.Threshould {
		service.mu.RUnlock()

		service.mu.Lock()
		service.state = states.Open
		service.mu.Unlock()

		c <- callResponse{nil, errs.ErrGrelayStateOpened}
		return
	}
	service.mu.RUnlock()

	i, err := f()
	c <- callResponse{i, err}
}

func (g grelayExecWithGo) processTimeout(service *Service) error {
	service.mu.Lock()
	service.currentServiceThreshould++
	service.mu.Unlock()

	service.mu.RLock()
	if service.currentServiceThreshould >= service.config.Threshould {
		service.mu.RUnlock()

		service.mu.Lock()
		service.state = states.Open
		service.mu.Unlock()

		return errs.ErrGrelayServiceTimedout
	}
	service.mu.RUnlock()

	return errs.ErrGrelayServiceTimedout
}
