package grelay

import (
	"sync"
	"testing"
	"time"

	"github.com/grelay/grelay/internal/errs"
	"github.com/grelay/grelay/internal/states"
	"github.com/stretchr/testify/assert"
)

func TestGrelayExecWithGoWithClosedState(t *testing.T) {
	c := DefaultConfiguration
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,

		mu: &sync.RWMutex{},
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.Nil(t, err, "Error Should return nil")
}

func TestGrelayExecWithGoWithOpenState(t *testing.T) {
	c := DefaultConfiguration
	g := &Service{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 0,

		mu: &sync.RWMutex{},
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, errs.ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithHalfOpenState(t *testing.T) {
	c := DefaultConfiguration
	g := &Service{
		config:                   c,
		state:                    states.HalfOpen,
		currentServiceThreshould: 0,

		mu: &sync.RWMutex{},
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.HalfOpen), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, errs.ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithClosedStateWithCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.Threshould = 5
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 6,

		mu: &sync.RWMutex{},
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(6), g.currentServiceThreshould)
	assert.EqualError(t, err, errs.ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldLessThanServiceThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.Threshould = 5
	c.Timeout = 5 * time.Microsecond
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,

		mu: &sync.RWMutex{},
	}

	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		time.Sleep(5 * time.Second)
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(4), g.currentServiceThreshould)
	assert.EqualError(t, err, errs.ErrGrelayServiceTimedout.Error())
}

func TestGrelayExecWithGoWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.Threshould = 5
	c.Timeout = 5 * time.Microsecond
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 4,

		mu: &sync.RWMutex{},
	}

	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		time.Sleep(5 * time.Second)
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(5), g.currentServiceThreshould)
	assert.EqualError(t, err, errs.ErrGrelayServiceTimedout.Error())
}
