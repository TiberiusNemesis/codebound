package data_structures_impl

import (
	"testing"
)

// TestPush tests the push functionality of the stack.
func TestStack_Push(t *testing.T) {
	stack := Stack{}

	// Test pushing valid data
	err := stack.push(10)
	if err != nil {
		t.Errorf("push returned an error: %v", err)
	}

	// Test pushing nil
	err = stack.push(nil)
	if err == nil {
		t.Errorf("Expected an error when pushing nil, but got nil")
	}
}

// TestPop tests the pop functionality of the stack.
func TestStack_Pop(t *testing.T) {
	stack := Stack{}
	stack.push(10)
	stack.push(20)

	// Test popping elements
	element, err := stack.pop()
	if err != nil {
		t.Errorf("pop returned an error: %v", err)
	}
	if element.data != 20 {
		t.Errorf("Expected 20, got %v", element.data)
	}

	// Test popping from an empty stack
	stack.pop() // Now the stack should be empty
	_, err = stack.pop()
	if err == nil {
		t.Errorf("Expected an error when popping from an empty stack, but got nil")
	}
}

// TestPeek tests the peek functionality of the stack.
func TestStack_Peek(t *testing.T) {
	stack := Stack{}
	stack.push(10)

	// Test peeking
	element, err := stack.peek()
	if err != nil {
		t.Errorf("peek returned an error: %v", err)
	}
	if element.data != 10 {
		t.Errorf("Expected 10, got %v", element.data)
	}

	// Test peeking at an empty stack
	stack.pop() // Empty the stack
	_, err = stack.peek()
	if err == nil {
		t.Errorf("Expected an error when peeking at an empty stack, but got nil")
	}
}

// Optionally, more tests can be added to check the integrity of the stack after various operations.
