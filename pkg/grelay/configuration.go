package grelay

import (
	"time"
)

// Pingable provides a Ping contract.
//
// This contract is a Ping function which has the responsability to check the health of the service.
// When Ping function returns nil, it means the service is healthy.
type Pingable interface {
	Ping() error
}

// Configuration is a structure for GrelayService configuration
//
// TODO remove the Service configuration.
type Configuration struct {
	// WithGo is responsible to execute in a goroutine your call function.
	// When WithGo is set to true, and the Timeout condiguration is reached, the request will return and cancel the request.
	// When WithGo is set to false, all request will wait even if the Timeout configuration is reached.
	WithGo bool
	// Threshould sets the limit of threshould. to change the state of the service to OPEN. It means that the
	// When Threshould hits the limit, the service will be blocked by request until it be healthy again.
	//
	// Grelay will check the healthy of the application by using the Service configuration.
	Threshould int64
	// RetryPeriod sets the retry time period to check the health of the service when it is blocked.
	// We don't recommend to put a short value (i.e. 1 microsecond), because it can generate a bottleneck in your application.
	RetryPeriod time.Duration
	// Timeout sets the limit of time that the service can take to increase the threshould.
	Timeout time.Duration
	// Service is responsible to check the health of an specific service (can be a microservice, database or any external/internal application).
	// Service Pingable
}

// DefaultConfiguration is a default configuration for a service.
var DefaultConfiguration = Configuration{
	RetryPeriod: 10 * time.Second,
	Timeout:     10 * time.Second,
	Threshould:  10,
	WithGo:      true,
}
