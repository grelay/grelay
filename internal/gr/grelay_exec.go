package gr

type grelayExec interface {
	exec(GrelayService, func() (interface{}, error)) (interface{}, error)
}
