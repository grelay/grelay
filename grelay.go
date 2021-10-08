package grelay

type Grelay interface {
	Exec(func() (interface{}, error)) (interface{}, error)
}

type grelayImpl struct {
	config GrelayConfig
}

func NewGrelay(c GrelayConfig) Grelay {
	return &grelayImpl{
		config: c,
	}
}

func (g *grelayImpl) Exec(f func() (interface{}, error)) (interface{}, error) {
	return f()
}
