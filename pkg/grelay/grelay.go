package grelay

import (
	"sync"

	"github.com/grelay/grelay/internal/gr"
)

type GrelayService struct {
	gs gr.GrelayService
}

// Grelay is an interface that have Enqueue function
type Grelay interface {
	/* CreateRequest is responsable to create a GrelayRequest to enqueue new functions later.

	EX:

	gr := grelay.CreateRequest()
	*/
	CreateRequest() gr.GrelayRequest
}

type grelayImpl struct {
	mapServices map[string]gr.GrelayService
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
	ms := make(map[string]gr.GrelayService, len(m))
	for k, v := range m {
		ms[k] = v.gs
	}
	return &grelayImpl{
		mapServices: ms,
	}
}

func (g *grelayImpl) CreateRequest() gr.GrelayRequest {
	return gr.GrelayRequestImpl{
		MapServices: g.mapServices,
		Mu:          &sync.RWMutex{},
	}
}

func NewGrelayConfig() gr.GrelayConfig {
	return gr.NewGrelayConfig()
}

func NewGrelayService(c gr.GrelayConfig) GrelayService {
	return GrelayService{
		gs: gr.NewGrelayService(c),
	}
}
