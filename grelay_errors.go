package grelay

import "errors"

// ErrGrelayServiceTimedout indicates that the service timed out
var ErrGrelayServiceTimedout = errors.New("gRelay error - service timed out")

// ErrGrelayStateOpened indicates that the service has open states
var ErrGrelayStateOpened = errors.New("gRelay error - state is opened")

// ErrGrelayAllRequestsOpened indicates that all services enqueued are open
var ErrGrelayAllRequestsOpened = errors.New("gRelay error - all requests are opened")
