package data_structures_impl

import (
	"fmt"
)

type Stack struct {
	elements []StackElement 
}

type StackElement struct {
	data any
}

type Stacker interface {
	pop() (StackElement, error)
	peek() (StackElement, error)	
	push(data any) error
}

// Removes the last element from the stack
func (stack *Stack) pop() (*StackElement, error) {
	lastElement, err := stack.peek()
	if err != nil {
		return nil, err
	}
	fmt.Println("Element being popped: ", lastElement)
	
	stack.elements = stack.elements[0:len(stack.elements)-1]
	fmt.Println("Stack after pop: ", stack.elements)
	return lastElement, nil
}

// Adds a new element to the end of the stack
func (stack *Stack) push(newElementData any) error {
	if (newElementData != nil) {
		newElement := StackElement{data: newElementData}
		stack.elements = append(stack.elements, newElement)
		return nil
	}
	return fmt.Errorf("Data to be added is null.")
}

// Returns the last element of the stack
func (stack *Stack) peek() (*StackElement, error) {
	if (len(stack.elements) > 0) {
		return &stack.elements[len(stack.elements)-1], nil
	}
	return nil, fmt.Errorf("There are no elements in the stack.")
}