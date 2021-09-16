package common

type ResizeStatus int

const (
	CREATED ResizeStatus = iota
	STARTED
	FINISHED
	ERROR
)
