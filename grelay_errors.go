package grelay

import "errors"

var ErrGrelayServiceTimedout = errors.New("gRelay error - service timed out")
var ErrGrelayStateOpened = errors.New("gRelay error - state is opened")
var ErrGrelayAllRequestsOpened = errors.New("gRelay error - all requests are opened")
