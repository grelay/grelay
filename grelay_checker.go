package grelay

// GrelayChecker has a ping function that is responsible for ping to the server when the state is OPEN
type GrelayChecker interface {
	Ping() error
}
