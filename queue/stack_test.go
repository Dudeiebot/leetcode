package main

import (
	"testing"
)

// func TestMyStackOperations(t *testing.T) {
// 	stack := Constructor()
//
// 	// Test Empty Stack
// 	if !stack.Empty() {
// 		t.Error("Expected stack to be empty")
// 	}
//
// 	// Test Push, Top, and Empty
// 	stack.Push(1)
// 	if top := stack.Top(); top != 1 {
// 		t.Errorf("Expected top to be 1, got %v", top)
// 	}
// 	if stack.Empty() {
// 		t.Error("Expected stack not to be empty")
// 	}
//
// 	// Test Push and Pop
// 	stack.Push(2)
// 	stack.Push(3)
// 	if pop := stack.Pop(); pop != 3 {
// 		t.Errorf("Expected pop to be 3, got %v", pop)
// 	}
// 	if top := stack.Top(); top != 2 {
// 		t.Errorf("Expected top to be 2, got %v", top)
// 	}
//
// 	// Test Empty after Pop
// 	stack.Pop()
// 	stack.Pop()
// 	if !stack.Empty() {
// 		t.Error("Expected stack to be empty after all pops")
// 	}
//
// 	// Test Top on Empty Stack
// 	top := stack.Top()
// 	if top != 0 {
// 		t.Errorf("Expected top to be 0 on an empty stack, got %v", top)
// 	}
// }

// func TestMyQueueOperations(t *testing.T) {
// 	queue := Constructor()
//
// 	// Test Empty Stack
// 	if !queue.Empty() {
// 		t.Error("Expected stack to be empty")
// 	}
//
// 	// Test Push, Top, and Empty
// 	queue.Push(1)
// 	if top := queue.Peek(); top != 1 {
// 		t.Errorf("Expected top to be 1, got %v", top)
// 	}
// 	if queue.Empty() {
// 		t.Error("Expected stack not to be empty")
// 	}
//
// 	// Test Push and Pop
// 	queue.Push(2)
// 	queue.Push(3)
// 	if pop := queue.Pop(); pop != 1 {
// 		t.Errorf("Expected pop to be 1, got %v", pop)
// 	}
// 	if top := queue.Peek(); top != 2 {
// 		t.Errorf("Expected top to be 2, got %v", top)
// 	}
//
// 	// Test Empty after Pop
// 	queue.Pop()
// 	queue.Pop()
// 	if !queue.Empty() {
// 		t.Error("Expected stack to be empty after all pops")
// 	}
//
// 	// Test Top on Empty Stack
// 	top := queue.Peek()
// 	if top != 0 {
// 		t.Errorf("Expected top to be 0 on an empty stack, got %v", top)
// 	}
// }

func TestMinStack(t *testing.T) {
	stack := Constructor()

	// Test Push and GetMin
	stack.Push(5)
	if stack.GetMin() != 5 {
		t.Errorf("Expected GetMin to be 5, got %d", stack.GetMin())
	}

	stack.Push(3)
	if stack.GetMin() != 3 {
		t.Errorf("Expected GetMin to be 3, got %d", stack.GetMin())
	}

	stack.Push(7)
	if stack.GetMin() != 3 {
		t.Errorf("Expected GetMin to still be 3, got %d", stack.GetMin())
	}

	stack.Push(2)
	if stack.GetMin() != 2 {
		t.Errorf("Expected GetMin to be 2, got %d", stack.GetMin())
	}

	// Test Pop and GetMin
	stack.Pop()
	if stack.GetMin() != 3 {
		t.Errorf("Expected GetMin to return to 3 after popping 2, got %d", stack.GetMin())
	}

	stack.Pop()
	if stack.GetMin() != 3 {
		t.Errorf("Expected GetMin to remain 3 after popping 7, got %d", stack.GetMin())
	}

	stack.Pop()
	if stack.GetMin() != 5 {
		t.Errorf("Expected GetMin to be 5 after popping 3, got %d", stack.GetMin())
	}

	// Test Top
	stack.Push(8)
	if stack.Top() != 8 {
		t.Errorf("Expected Top to be 8, got %d", stack.Top())
	}

	stack.Pop()
	if stack.Top() != 5 {
		t.Errorf("Expected Top to be 5 after popping 8, got %d", stack.Top())
	}
}
