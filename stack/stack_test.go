package stack

import "testing"

func TestEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Fatalf("expected empty stack")
	}

	got := s.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestNotEmpty(t *testing.T) {
	s := New[int]()
	s.Push(1)

	if s.IsEmpty() {
		t.Fatalf("expected non-empty stack")
	}
	got := s.Size()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPushPop(t *testing.T) {
	s := New[int]()
	//check size
	got := s.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	//check empty
	if !s.IsEmpty() {
		t.Fatalf("expected empty stack")
	}

	//push 3 values
	s.Push(3)
	s.Push(2)
	s.Push(1)

	//check size
	got = s.Size()
	if want := 3; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//pop one
	got = s.Pop()
	if want := 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = s.Size()
	if want := 2; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//pop one
	got = s.Pop()
	if want := 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = s.Size()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	//pop one
	got = s.Pop()
	if want := 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	//check size
	got = s.Size()
	if want := 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	//check empty
	if !s.IsEmpty() {
		t.Fatalf("expected empty stack")
	}
}

func TestPopEmpty(t *testing.T) {
	s := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected Pop to panic")
		}
	}()

	s.Pop()
}

func TestPeek(t *testing.T) {
	s := New[int]()
	s.Push(1)

	if s.IsEmpty() {
		t.Fatalf("expected non-empty stack")
	}
	//peek value
	got := s.Peek()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	//test non-empty
	if s.IsEmpty() {
		t.Fatalf("expected non-empty stack")
	}
	//test value not removed
	got = s.Peek()
	if want := 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
