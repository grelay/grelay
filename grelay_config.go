package grelay

import "time"

type GrelayConfig struct {
	retryTimePeriod   time.Duration
	serviceTimeout    time.Duration
	serviceThreshould int64
	service           GrelayChecker
}

/* NewGrelayConfig create config for grelay with these values:

- RetryTimePeriod: 10s

- ServiveTimeout: 10s

- ServiceThreshould: 10 times
*/
func NewGrelayConfig() GrelayConfig {
	return GrelayConfig{
		retryTimePeriod:   10 * time.Second,
		serviceTimeout:    10 * time.Second,
		serviceThreshould: 10,
		service:           nil,
	}
}

/* WithRetryTimePeriod sets the retry time period when the state is OPEN.

ATTENTION: Do not put a really short time (EX: 1 microsecond) to not lock a lot your application
*/
func (c GrelayConfig) WithRetryTimePeriod(t time.Duration) GrelayConfig {
	c.retryTimePeriod = t
	return c
}

// WithServiceTimeout sets the limit of time that the service can take to increase the threashould
func (c GrelayConfig) WithServiceTimeout(t time.Duration) GrelayConfig {
	c.serviceTimeout = t
	return c
}

// WithServiceThreshould sets the limit of threshould to change the state to OPEN
func (c GrelayConfig) WithServiceThreshould(ts int64) GrelayConfig {
	c.serviceThreshould = ts
	return c
}

// WithGrelayService sets the service that is responsible for ping to the server when the state is OPEN
func (c GrelayConfig) WithGrelayService(service GrelayChecker) GrelayConfig {
	c.service = service
	return c
}
