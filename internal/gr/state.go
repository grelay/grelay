package gr

type state string

const (
	closed   state = "CLOSED"
	open     state = "OPEN"
	halfOpen state = "HALF-OPEN"
)
