package grelay

import "sync"

type Grelay interface {
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
