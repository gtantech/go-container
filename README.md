# go-container
A Go package that provides a generic implementation of a queue and stack for FIFO and LIFO operations.

# Install
Install via `go get`. Note that Go 1.25 or newer is required.

```sh
# After: go mod init ...
go get -u github.com/gtantech/go-container
```

# Example

```go
package main

import (
	"fmt"

	"github.com/gtantech/go-container/queue"
	"github.com/gtantech/go-container/stack"
)

func main() {
	s := stack.New[int]()
	q := queue.New[int]()

	fmt.Printf("Stack is empty: %v, with size: %v\n", s.IsEmpty(), s.Size())
	fmt.Printf("Queue is empty: %v, with size: %v\n", q.IsEmpty(), s.Size())

	s.Push(1)
	s.Push(2)
	s.Push(3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Printf("Stack is empty: %v, with size: %v\n", s.IsEmpty(), s.Size())
	fmt.Printf("Queue is empty: %v, with size: %v\n", q.IsEmpty(), s.Size())

	fmt.Printf("Stack peeked: %v\n", s.Peek())
	fmt.Printf("Stack popped: %v\n", s.Pop())
	fmt.Printf("Stack popped: %v\n", s.Pop())
	fmt.Printf("Stack popped: %v\n", s.Pop())

	fmt.Printf("Queue dequeued: %v\n", q.Dequeue())
	fmt.Printf("Queue dequeued: %v\n", q.Dequeue())
	fmt.Printf("Queue dequeued: %v\n", q.Dequeue())
}

```