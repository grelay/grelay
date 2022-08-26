package grelay

import (
	"time"
)

// Configuration is a structure for GrelayService configuration
type Configuration struct {
	withGo            bool
	serviceThreshould int64
	retryTimePeriod   time.Duration
	serviceTimeout    time.Duration
	service           GrelayChecker
}

/* NewGrelayConfig create config for grelay with these values:

- RetryTimePeriod: 10s

- ServiveTimeout: 10s

- ServiceThreshould: 10 times
*/
func NewGrelayConfig() Configuration {
	return Configuration{
		retryTimePeriod:   10 * time.Second,
		serviceTimeout:    10 * time.Second,
		serviceThreshould: 10,
		service:           nil,
		withGo:            true,
	}
}

/* WithRetryTimePeriod sets the retry time period when the state is OPEN.

ATTENTION: Do not put a really short time (EX: 1 microsecond) to not lock a lot your application
*/
func (c Configuration) WithRetryTimePeriod(t time.Duration) Configuration {
	c.retryTimePeriod = t
	return c
}

// WithServiceTimeout sets the limit of time that the service can take to increase the threashould
func (c Configuration) WithServiceTimeout(t time.Duration) Configuration {
	c.serviceTimeout = t
	return c
}

// WithServiceThreshould sets the limit of threshould to change the state to OPEN
func (c Configuration) WithServiceThreshould(ts int64) Configuration {
	c.serviceThreshould = ts
	return c
}

// WithGrelayService sets the service that is responsible for ping to the server when the state is OPEN
func (c Configuration) WithGrelayService(service GrelayChecker) Configuration {
	c.service = service
	return c
}

// WithGo is responsible to execute in a goroutine your call function. If get ServiceTimeout, grelay will return instead of hold your call.
func (c Configuration) WithGo() Configuration {
	c.withGo = true
	return c
}
