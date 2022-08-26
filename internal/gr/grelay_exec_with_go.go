package gr

import "time"

type grelayExecWithGo struct{}

func (e grelayExecWithGo) exec(service GrelayService, f func() (interface{}, error)) (interface{}, error) {
	gs := service.(*grelayServiceImpl)
	callDone := make(chan callResponse, 1)
	go e.makeCall(gs, f, callDone)

	gs.mu.RLock()
	t := time.NewTimer(gs.config.serviceTimeout)
	gs.mu.RUnlock()
	defer t.Stop()

	select {
	case <-t.C:
		gs.mu.Lock()
		gs.currentServiceThreshould++
		gs.mu.Unlock()

		gs.mu.RLock()
		if gs.currentServiceThreshould >= gs.config.serviceThreshould {
			gs.mu.RUnlock()

			gs.mu.Lock()
			gs.state = open
			gs.mu.Unlock()

			return nil, ErrGrelayServiceTimedout
		}
		gs.mu.RUnlock()

		return nil, ErrGrelayServiceTimedout
	case r := <-callDone:
		return r.i, r.err
	}
}

func (g grelayExecWithGo) makeCall(service GrelayService, f func() (interface{}, error), c chan<- callResponse) {
	gs := service.(*grelayServiceImpl)
	defer close(c)
	gs.mu.RLock()
	if string(gs.state) == string(open) || string(gs.state) == string(halfOpen) {
		gs.mu.RUnlock()
		c <- callResponse{nil, ErrGrelayStateOpened}
		return
	}
	if gs.currentServiceThreshould >= gs.config.serviceThreshould {
		gs.mu.RUnlock()

		gs.mu.Lock()
		gs.state = open
		gs.mu.Unlock()

		c <- callResponse{nil, ErrGrelayStateOpened}
		return
	}
	gs.mu.RUnlock()

	i, err := f()
	c <- callResponse{i, err}
}
