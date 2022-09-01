package grelay

import (
	"errors"
	"sync"
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

func createGrelayService(t time.Duration, err error) Pingable {
	return &grelayServiceTest{
		t:   t,
		err: err,
	}
}

func TestExecWithGo(t *testing.T) {
	c := DefaultConfiguration
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,

		mu: &sync.RWMutex{},
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
	c := DefaultConfiguration
	c.RetryPeriod = 5 * time.Microsecond

	s := createGrelayService(1*time.Microsecond, nil)
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}

	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndCurrentServiceThreshouldEqualZeroShouldKeepThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.RetryPeriod = 5 * time.Microsecond

	s := createGrelayService(1*time.Microsecond, nil)
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 0,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndPingServiceTimedoutShouldKeepThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.RetryPeriod = 10 * time.Millisecond
	c.Timeout = 10 * time.Millisecond

	s := createGrelayService(50*time.Millisecond, nil)

	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateClosedAndServiceReturningErrorShouldKeepThreshould(t *testing.T) {
	c := DefaultConfiguration
	c.RetryPeriod = 5 * time.Microsecond

	s := createGrelayService(4*time.Microsecond, errors.New("Ping error"))
	g := &Service{
		config:                   c,
		state:                    states.Closed,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndPingSuccedShouldHaveClosedState(t *testing.T) {
	s := createGrelayService(1*time.Microsecond, nil)

	c := DefaultConfiguration

	g := &Service{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()
	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Closed), string(g.state))
	assert.Equal(t, int64(0), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndPingAndTimeoutDoesNotHaveTimeToAnswerShouldHaveHalfOpenStates(t *testing.T) {
	s := createGrelayService(5*time.Second, nil)
	c := DefaultConfiguration

	g := &Service{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
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

	c := DefaultConfiguration

	g := &Service{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}

func TestMonitoringWhenStateOpenAndTimeoutOccurredShouldHaveOpenState(t *testing.T) {
	s := createGrelayService(1*time.Second, nil)

	c := DefaultConfiguration
	c.Timeout = 5 * time.Microsecond

	g := &Service{
		config:                   c,
		state:                    states.Open,
		currentServiceThreshould: 3,
		service:                  s,

		mu: &sync.RWMutex{},
	}
	g.monitoringState()

	g.mu.RLock()
	defer g.mu.RUnlock()
	assert.Equal(t, string(states.Open), string(g.state))
	assert.Equal(t, int64(3), g.currentServiceThreshould)
}
