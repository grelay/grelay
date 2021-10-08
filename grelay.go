package grelay

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

func (g *grelayImpl) Enqueue(string, func() (interface{}, error)) GrelayRequest {
	return grelayRequestImpl{
		mapServices: g.mapServices,
	}
}
