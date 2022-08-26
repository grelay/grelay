package grelay

import (
	"errors"
	"testing"
	"time"

	"github.com/grelay/grelay/internal/states"
	"github.com/stretchr/testify/assert"
)

type grelayServiceTest struct {
	t   time.Duration
	err error
}

func (g *grelayServiceTest) Ping() error {
	time.Sleep(g.t)
	return g.err
}

func createGrelayService(t time.Duration, err error) GrelayChecker {
	return &grelayServiceTest{
		t:   t,
		err: err,
	}
}

func TestExecWithGo(t *testing.T) {
	c := NewGrelayConfig()
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,
	}
	_, err := g.exec(func() (interface{}, error) {
		return nil, nil
	})

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
	assert.Nil(t, err, "Error Should return nil")
}

func TestMonitoringWhenStateClosedAndServiceOKShouldResetThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithRetryTimePeriod(5 * time.Microsecond)

	s := createGrelayService(1*time.Microsecond, nil)
	c = c.WithGrelayService(s)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
	}

	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndCurrentServiceThreshouldEqualZeroShouldKeepThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithRetryTimePeriod(5 * time.Microsecond)

	s := createGrelayService(1*time.Microsecond, nil)
	c = c.WithGrelayService(s)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndServiceNotOKShouldKeepThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithRetryTimePeriod(7 * time.Millisecond)

	s := createGrelayService(10*time.Millisecond, nil)
	c = c.WithGrelayService(s)
	c = c.WithServiceTimeout(2 * time.Millisecond)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndServiceReturningErrorShouldKeepThreshould(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithRetryTimePeriod(5 * time.Microsecond)

	s := createGrelayService(4*time.Microsecond, errors.New("Ping error"))
	c = c.WithGrelayService(s)
	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndPingSuccedShouldHaveClosedState(t *testing.T) {
	s := createGrelayService(1*time.Microsecond, nil)

	c := NewGrelayConfig()
	c = c.WithGrelayService(s)

	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndPingAndTimeoutDoesNotHaveTimeToAnswerShouldHaveHalfOpenStates(t *testing.T) {
	s := createGrelayService(5*time.Second, nil)
	c := NewGrelayConfig()
	c = c.WithGrelayService(s)

	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
	}
	go g.monitoringState()
	time.Sleep(5 * time.Millisecond)
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.HalfOpen), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndPingFailedShouldHaveOpenState(t *testing.T) {
	s := createGrelayService(1*time.Microsecond, errors.New("Ping fail"))

	c := NewGrelayConfig()
	c = c.WithGrelayService(s)

	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndTimeoutOccurredShouldHaveOpenState(t *testing.T) {
	s := createGrelayService(1*time.Second, nil)

	c := NewGrelayConfig()
	c = c.WithGrelayService(s)
	c = c.WithServiceTimeout(5 * time.Microsecond)

	g := &grelayServiceImpl{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}
