package grelay

type grelayExec interface {
	exec(GrelayService, func() (interface{}, error)) (interface{}, error)
}
