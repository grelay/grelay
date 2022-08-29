package grelay

type grelayExec interface {
	exec(*Service, func() (interface{}, error)) (interface{}, error)
}
