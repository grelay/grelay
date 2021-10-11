package grelay

import "sync"

// Grelay is an interface that have Enqueue function
type Grelay interface {
	/* CreateRequest is responsable to create a GrelayRequest to enqueue new functions later.

	EX:

	gr := grelay.CreateRequest()
	*/
	CreateRequest() GrelayRequest
}

type grelayImpl struct {
	mapServices map[string]GrelayService
}

/* NewGrelay creates a grelay config using a map of string:GrelayService

EX:

	config1 := grelay.NewGrelayConfig()
	service1 := grelay.NewGrelayService(config1)

	config2 := grelay.NewGrelayConfig()
	service2 := grelay.NewGrelayService(config2)

	m := map[string]GrelayService{
		"service1": service1,
		"service2": service2,
	}

	g := NewGrelay(m)

*/
func NewGrelay(m map[string]GrelayService) Grelay {
	return &grelayImpl{
		mapServices: m,
	}
}

func (g *grelayImpl) CreateRequest() GrelayRequest {
	return grelayRequestImpl{
		mapServices: g.mapServices,
		mu:          &sync.RWMutex{},
	}
}
