package grelay

import "sync"

// Grelay is an interface that have Enqueue function
type Grelay interface {
	/* Enqueue is responsable to create a GrelayRequest and enqueue functions from a specific service in it.

	EX:

	gr := grelay.Enqueue("mygrelayservice", func() (interface{}, error) {
		value, err := myservice.GET("userID")
		if err != nil {
			return nil, err
		}
		return value, err
	})
	*/
	Enqueue(string, func() (interface{}, error)) GrelayRequest
}

type grelayImpl struct {
	mapServices map[string]GrelayService
}

func NewGrelay(m map[string]GrelayService) Grelay {
	return &grelayImpl{
		mapServices: m,
	}
}

func (g *grelayImpl) Enqueue(s string, f func() (interface{}, error)) GrelayRequest {
	gr := grelayRequestImpl{
		mapServices: g.mapServices,
		mu:          &sync.RWMutex{},
	}
	grEnqueued := gr.Enqueue(s, f)
	return grEnqueued
}
