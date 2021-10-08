package grelay

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecWithClosedState(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayImpl{
		config:                   c,
		state:                    closed,
		currentServiceThreshould: 0,
	}
	_, err := g.Exec(func() (interface{}, error) {
		return nil, nil
	})
	assert.Equal(t, string(closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.Nil(t, err, "Error Should return nil")
}

func TestExecWithOpenState(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayImpl{
		config:                   c,
		state:                    open,
		currentServiceThreshould: 0,
	}
	_, err := g.Exec(func() (interface{}, error) {
		return nil, nil
	})
	assert.Equal(t, string(open), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestExecWithHalfOpenState(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayImpl{
		config:                   c,
		state:                    halfOpen,
		currentServiceThreshould: 0,
	}
	_, err := g.Exec(func() (interface{}, error) {
		return nil, nil
	})
	assert.Equal(t, string(halfOpen), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestExecWithClosedStateWithCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	g := &grelayImpl{
		config:                   c,
		state:                    closed,
		currentServiceThreshould: 6,
	}
	_, err := g.Exec(func() (interface{}, error) {
		return nil, nil
	})
	assert.Equal(t, string(open), string(g.state))
	assert.Equal(t, int64(6), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayStateOpened.Error())
}

func TestExecWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldLessThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	c = c.WithServiceTimeout(5 * time.Microsecond)
	g := &grelayImpl{
		config:                   c,
		state:                    closed,
		currentServiceThreshould: 3,
	}

	_, err := g.Exec(func() (interface{}, error) {
		time.Sleep(5 * time.Second)
		return nil, nil
	})

	assert.Equal(t, string(closed), string(g.state))
	assert.Equal(t, int64(4), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayServiceTimedout.Error())
}

func TestExecWithClosedStateWithServiceTimeoutAndCurrentServiceThreshouldGratherThanServiceThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(5)
	c = c.WithServiceTimeout(5 * time.Microsecond)
	g := &grelayImpl{
		config:                   c,
		state:                    closed,
		currentServiceThreshould: 4,
	}

	_, err := g.Exec(func() (interface{}, error) {
		time.Sleep(5 * time.Second)
		return nil, nil
	})

	assert.Equal(t, string(open), string(g.state))
	assert.Equal(t, int64(5), g.currentServiceThreshould)
	assert.EqualError(t, err, ErrGrelayServiceTimedout.Error())
}
