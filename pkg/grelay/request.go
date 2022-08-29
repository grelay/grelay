package grelay

import (
	"errors"
	"fmt"
	"sync"

	"github.com/grelay/grelay/internal/errs"
	"go.uber.org/zap"
)

// GrelayRequest is an interface responsable to Enqueue and Exec requests
type GrelayRequest interface {
	/* Enqueue is responsable to enqueue functions from a specific service in GrelayRequest

	EX:

	gr := grelayRequest.Enqueue("mygrelayservice", func() (interface{}, error) {
		value, err := myservice.GET("userID")
		if err != nil {
			return nil, err
		}
		return value, err
	})
	*/
	Enqueue(string, func() (interface{}, error)) GrelayRequest

	/* Exec is responsable to execute requests from GrelayRequest queue.

	EX:

	gr := grelayRequest.Enqueue("mygrelayservice", func() (interface{}, error) {
		// make request1
	})
	gr = grelayRequest.Enqueue("mygrelayservice2", func() (interface{}, error) {
		// make request2
	})
	val, err := gr.Exec()
	*/
	Exec() (interface{}, error)
}

type GrelayRequestFunc func() (interface{}, error)

type GrelayRequestImpl struct {
	MapServices map[string]*Service
	QueueFuncs  []GrelayRequestFunc

	Mu *sync.RWMutex
}

func (gr GrelayRequestImpl) Enqueue(s string, f func() (interface{}, error)) GrelayRequest {
	gr.Mu.RLock()
	service, ok := gr.MapServices[s]
	gr.Mu.RUnlock()

	// TODO Remove zap and return an error
	if !ok {
		logger, _ := zap.NewProduction()
		sugar := logger.Sugar()
		gr.Mu.RLock()
		defer gr.Mu.RUnlock()
		sugar.Warn(fmt.Sprintf("grelay not found service with %s key", s), zap.Any("grelay_services", gr.MapServices))
		return gr
	}

	gr.Mu.Lock()
	if gr.QueueFuncs == nil {
		gr.QueueFuncs = []GrelayRequestFunc{}
	}
	gr.QueueFuncs = append(gr.QueueFuncs, func() (interface{}, error) {
		return service.exec(f)
	})
	gr.Mu.Unlock()

	gr.Mu.RLock()
	defer gr.Mu.RUnlock()
	return gr
}

func (gr2 GrelayRequestImpl) Exec() (interface{}, error) {
	gr2.Mu.RLock()
	defer gr2.Mu.RUnlock()
	for _, f := range gr2.QueueFuncs {
		value, err := f()
		if errors.Is(err, errs.ErrGrelayStateOpened) {
			continue
		}
		return value, err
	}
	return nil, errs.ErrGrelayAllRequestsOpened
}
