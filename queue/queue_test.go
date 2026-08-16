package queue

import "testing"

func TestEmpty(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Fatalf("expected empty queue")
	}

	got := q.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestNotEmpty(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)

	if q.IsEmpty() {
		t.Fatalf("expected non-empty queue")
	}
	got := q.Size()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := New[int]()
	//check size
	got := q.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	//check empty
	if !q.IsEmpty() {
		t.Fatalf("expected empty queue")
	}

	//enqueue 3 values
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	//check size
	got = q.Size()
	if want := 3; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//dequeue one
	got = q.Dequeue()
	if want := 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = q.Size()
	if want := 2; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//dequeue one
	got = q.Dequeue()
	if want := 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = q.Size()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//dequeue one
	got = q.Dequeue()
	if want := 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = q.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	//check empty
	if !q.IsEmpty() {
		t.Fatalf("expected empty queue")
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected Dequeue to panic")
		}
	}()

	q.Dequeue()
}
