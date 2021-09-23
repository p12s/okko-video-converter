package common

type ProcessStatus int

const (
	CREATED ProcessStatus = iota
	STARTED
	FINISHED
	ERROR
)
