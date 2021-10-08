package grelay

type Grelay interface {
	Exec(func() (interface{}, error)) (interface{}, error)
}

type grelayImpl struct {
}

func NewGrelay() Grelay {
	return &grelayImpl{}
}

func (g *grelayImpl) Exec(f func() (interface{}, error)) (interface{}, error) {
	return f()
}
