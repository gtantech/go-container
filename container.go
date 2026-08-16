package gocontainer

type Container interface {
	IsEmpty() bool
	Size() int
}
