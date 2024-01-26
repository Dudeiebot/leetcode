package main

import (
	"testing"
)

func TestMyStackOperations(t *testing.T) {
	stack := Constructor()

	// Test Empty Stack
	if !stack.Empty() {
		t.Error("Expected stack to be empty")
	}

	// Test Push, Top, and Empty
	stack.Push(1)
	if top := stack.Top(); top != 1 {
		t.Errorf("Expected top to be 1, got %v", top)
	}
	if stack.Empty() {
		t.Error("Expected stack not to be empty")
	}

	// Test Push and Pop
	stack.Push(2)
	stack.Push(3)
	if pop := stack.Pop(); pop != 3 {
		t.Errorf("Expected pop to be 3, got %v", pop)
	}
	if top := stack.Top(); top != 2 {
		t.Errorf("Expected top to be 2, got %v", top)
	}

	// Test Empty after Pop
	stack.Pop()
	stack.Pop()
	if !stack.Empty() {
		t.Error("Expected stack to be empty after all pops")
	}

	// Test Top on Empty Stack
	top := stack.Top()
	if top != 0 {
		t.Errorf("Expected top to be 0 on an empty stack, got %v", top)
	}
}
