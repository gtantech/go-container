package container

type Container interface {
	IsEmpty() bool
	Size() int
}
