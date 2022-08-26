package gr

import (
	"testing"
	"time"

	"github.com/grelay/grelay/internal/states"
	"github.com/stretchr/testify/assert"
)

func TestGrelayExecWithGoWithClosedState(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,
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
	c := NewGrelayConfig()
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 0,
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithHalfOpenState(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.HalfOpen,
		currentServiceThreshould: 0,
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.HalfOpen), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithClosedStateWithCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 6,
	}
	gExec := grelayExecWithGo{}
	_, err := gExec.exec(g, func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(6), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestGrelayExecWithGoWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldLessThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	c = c.WithServiceTimeout(5 * time.Microsecond)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
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
	assert.EqualError(t, err, ErrGrelayServiceTimedout.Error())
}

func TestGrelayExecWithGoWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	c = c.WithServiceTimeout(5 * time.Microsecond)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 4,
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
	assert.EqualError(t, err, ErrGrelayServiceTimedout.Error())
}
