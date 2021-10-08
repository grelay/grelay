package grelay

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type GrelayRequest interface {
	Enqueue(string, func() (interface{}, error)) GrelayRequest
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
		value, err := r.service.Exec(r.f)
		if errors.Is(err, ErrGrelayStateOpened) {
			continue
		}
		return value, err
	}
	return nil, ErrGrelayAllRequestsOpened
}
