package grelay

import (
	"errors"
	"fmt"
	"sync"

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

type grelayRequestQueueStruct struct {
	service GrelayService
	f       func() (interface{}, error)
}

type grelayRequestImpl struct {
	mapServices map[string]GrelayService
	QueueFuncs  []grelayRequestQueueStruct

	mu *sync.RWMutex
}

func (gr grelayRequestImpl) Enqueue(s string, f func() (interface{}, error)) GrelayRequest {
	gr.mu.RLock()
	service, ok := gr.mapServices[s]
	gr.mu.RUnlock()

	if !ok {
		logger, _ := zap.NewProduction()
		sugar := logger.Sugar()
		gr.mu.RLock()
		defer gr.mu.RUnlock()
		sugar.Warn(fmt.Sprintf("grelay no found service with %s key", s), zap.Any("grelay_services", gr.mapServices))
		return gr
	}

	gr.mu.Lock()
	if gr.QueueFuncs == nil {
		gr.QueueFuncs = []grelayRequestQueueStruct{}
	}
	gr.QueueFuncs = append(gr.QueueFuncs, grelayRequestQueueStruct{
		service: service,
		f:       f,
	})
	gr.mu.Unlock()

	gr.mu.RLock()
	defer gr.mu.RUnlock()
	return gr
}

func (gr grelayRequestImpl) Exec() (interface{}, error) {
	gr.mu.RLock()
	defer gr.mu.RUnlock()
	for _, r := range gr.QueueFuncs {
		value, err := r.service.exec(r.f)
		if errors.Is(err, ErrGrelayStateOpened) {
			continue
		}
		return value, err
	}
	return nil, ErrGrelayAllRequestsOpened
}
