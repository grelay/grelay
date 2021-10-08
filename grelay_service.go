package grelay

// GrelayService has a ping function that is responsible for ping to the server when the state is OPEN
type GrelayService interface {
	Ping() error
}
