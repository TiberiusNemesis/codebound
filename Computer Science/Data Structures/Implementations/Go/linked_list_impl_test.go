package data_structures_impl

import (
	"testing"
)

func TestLinkedList_GetNode(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// Test for getting nodes at valid indices
	expectedNode := firstNode
	actualNode, err := list.getNode(0)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 0: %v", err)
	}
	if actualNode != expectedNode {
		t.Errorf("Node at index 0 should've been %v, but was %v", expectedNode, actualNode)
	}

	secondExpectedNode := secondNode
	secondActualNode, err := list.getNode(1)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 1: %v", err)
	}
	if secondActualNode != secondExpectedNode {
		t.Errorf("Expected node at index 1 to be %v, but was %v", secondExpectedNode, secondActualNode)
	}

	thirdExpectedNode := thirdNode
	thirdActualNode, err := list.getNode(2)
	if err != nil {
		t.Errorf("Error occurred while getting node at index 2: %v", err)
	}
	if thirdActualNode != thirdExpectedNode {
		t.Errorf("Expected node at index 2 to be %v, but was %v", thirdExpectedNode, thirdActualNode)
	}

	// Test to get node at an invalid index
	_, err = list.getNode(3)
	if err == nil {
		t.Error("Expected error when getting node at index 3, but no error was thrown")
	}
}

func TestLinkedList_GetNode_InvalidIndex(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// Test to get node at an invalid index
	_, err := list.getNode(3)
	if err == nil {
		t.Error("Expected error when getting node at index 3, but no error was thrown")
	}
}

func TestLinkedList_GetNode_EmptyList(t *testing.T) {
	list := LinkedList{}

	// Test to get node at an invalid index
	_, err := list.getNode(0)
	if err == nil {
		t.Error("Expected error when getting node at index 0, but no error was thrown")
	}
}

func TestLinkedList_GetNode_IndexLowerThanZero(t *testing.T) {
	list := LinkedList{}

	// Test to get node at an index below zero
	_, err := list.getNode(-1)
	if err == nil {
		t.Error("Expected error when getting node at index -1, but no error was thrown")
	}
}

func TestLinkedList_RemoveNode(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// We remove the second node 
	err := list.removeNode(secondNode)
	if err != nil {
		t.Errorf("Error occurred while removing node: %v", err)
	}

	// Then verify that the list no longer has 3 indexes
	_, err = list.getNode(2)
	if err == nil {
		t.Error("Expected an error after removing when trying to get node at index 2, but no error was thrown")
	}
}

func TestLinkedList_RemoveNode_OnlyNode(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	list.appendNode(firstNode)

	// We remove the first node
	err := list.removeNode(firstNode)
	if err != nil {
		t.Errorf("Error occurred while removing node: %v", err)
	}

	// Then we verify that the list no longer has 3 indexes
	_, err = list.getNode(0)
	if err == nil {
		t.Error("Expected an error after removing when trying to get node at index 0, but no error was thrown")
	}
}

func TestLinkedList_RemoveNode_Head(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)

	// We remove the first node
	err := list.removeNode(firstNode)
	if err != nil {
		t.Errorf("Error occurred while removing head node: %v", err)
	}

	// Then we verify that the list no longer has 3 indexes
	_, err = list.getNode(2)
	if err == nil {
		t.Error("Expected an error after removing when trying to get node at index 2, but no error was thrown")
	}
}

func TestLinkedList_RemoveNode_Tail(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)

	// We remove the second node
	err := list.removeNode(secondNode)
	if err != nil {
		t.Errorf("Error occurred while removing node: %v", err)
	}

	// Then we verify that the list no longer has 3 indexes
	_, err = list.getNode(1)
	if err == nil {
		t.Error("Expected an error after removing when trying to get node at index 1, but no error was thrown")
	}
}

func TestLinkedList_ReplaceNode(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// We replace the second node with a new one
	newNode := &ListNode{data: "dark side"}
	err := list.replaceNode(1, newNode)
	if err != nil {
		t.Errorf("Error occurred while replacing node: %v", err)
	}

	// Then we verify that the second node really has been replaced
	replacedNode, err := list.getNode(1)
	if err != nil {
		t.Errorf("Error occurred while getting the replaced node: %v", err)
	}
	if replacedNode != newNode {
		t.Errorf("Expected replaced node to be %v, but was %v", newNode, replacedNode)
	}
}

// Write a test case for ReplaceNode when the index for the node to be replaced points to a non-existent node

func TestLinkedList_ReplaceNode_InvalidIndex(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)
	list.appendNode(thirdNode)

	// We replace the second node with a new one
	newNode := &ListNode{data: "dark side"}
	err := list.replaceNode(3, newNode)
	if err == nil {
		t.Error("Expected error when replacing node at index 3, but no error was thrown")
	}
}

func TestLinkedList_ReplaceNode_SingleNodeList(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "lonely"}
	list.appendNode(firstNode)
	
	// We replace the first node with a new one
	newNode := &ListNode{data: "akon"}
	err := list.replaceNode(0, newNode)
	if err != nil {
		t.Errorf("Error occurred while replacing node at head.")
	}

	// Then we verify that the tail/head node really has been replaced
	replacedNode, err := list.getNode(0)
	if err != nil {
		t.Errorf("Error occurred while getting the replaced node: %v", err)
	}
	if replacedNode != newNode {
		t.Errorf("Expected replaced node to be %v, but was %v", newNode, replacedNode)
	}
}

func TestLinkedList_AddNode_AtHead(t *testing.T) {
	list := LinkedList{}
	firstNode := &ListNode{data: "use"}
	secondNode := &ListNode{data: "the"}
	thirdNode := &ListNode{data: "force"}
	list.appendNode(firstNode)
	list.appendNode(secondNode)	
	list.appendNode(thirdNode)

	
	// We add a node at index 0 (the head)
	newNode := &ListNode{data: "just"}

	err := list.addNode(0, newNode)
	if err != nil {
		t.Errorf("Error occurred while adding node: %v", err)
	}

	// Then verify that the it was added at the correct index
	node, err := list.getNode(0)
	if err != nil {
		t.Errorf("Error occurred while getting the added node: %v", err)
	}
	if node != newNode {
		t.Errorf("Expected added node to be %v, but was %v", secondNode, node)
	}
}

func TestLinkedList_AppendNode(t *testing.T) {
	list := LinkedList{}
	node := &ListNode{data: "test data"} // I'm getting out of ideas here for test data (not that I had many to begin with)

	list.appendNode(node)

	if list.head != node {
		t.Errorf("Expected head to be %v, but got %v", node, list.head)
	}

	if list.tail != node {
		t.Errorf("Expected tail to be %v, but got %v", node, list.tail)
	}
}

func TestListNode_GetData(t *testing.T) {
	node := &ListNode{data: "I am the senate"}

	data, err := node.getData()

	if err != nil {
		t.Errorf("Error occurred while getting data: %v", err)
	}

	expectedData := "I am the senate"
	if data != expectedData {
		t.Errorf("Expected data to be %v, but got %v", expectedData, data)
	}
}

func TestListNode_GetData_NodeIsEmpty(t *testing.T) {
	node := &ListNode{} // Create an empty node

	data, err := node.getData()

	if err == nil {
		t.Errorf("Expected error to occur, but got nil")
	}

	if data != nil {
		t.Errorf("Expected data to be nil, but was %v", data)
	}
}

func TestListNode_GetPreviousNode(t *testing.T) {
	// Create a previous node
	previousNode := &ListNode{data: "Previous Node"}

	// Create a current node
	currentNode := &ListNode{data: "Current Node", previousPointer: previousNode}

	// Call the getPreviousNode method
	previous, err := currentNode.getPreviousNode()

	// Check if an error occurred
	if err != nil {
		t.Errorf("Error occurred while getting previous node: %v", err)
	}

	// Check if the previous node is correct
	if previous != previousNode {
		t.Errorf("Expected previous node to be %v, but got %v", previousNode, previous)
	}
}

func TestListNode_GetPreviousNode_NoPreviousNodeAvailable(t *testing.T) {
	// Create a current node with no previous node
	currentNode := &ListNode{data: "Current Node"}

	// Call the getPreviousNode method
	previous, err := currentNode.getPreviousNode()

	// Check if an error occurred
	if err == nil {
		t.Errorf("Error did not occur while getting previous node.")
	}

	// Check if the previous node is nil
	if previous != nil {
		t.Errorf("Expected previous node to be nil, but was %v", previous)
	}
}

func TestListNode_GetNextNode(t *testing.T) {
	// Create a next node
	nextNode := &ListNode{data: "Next Node"}

	// Create a current node
	currentNode := &ListNode{data: "Current Node", nextPointer: nextNode}

	// Call the getNextNode method
	next, err := currentNode.getNextNode()

	// Check if an error occurred
	if err != nil {
		t.Errorf("Error occurred while getting next node: %v", err)
	}

	// Check if the next node is correct
	if next != nextNode {
		t.Errorf("Expected next node to be %v, but got %v", nextNode, next)
	}
}

func TestLinkNodes(t *testing.T) {
	previousNode := &ListNode{data: "Previous Node"}
	newNode := &ListNode{data: "New Node"}
	nextNode := &ListNode{data: "Next Node"}

	// Call the linkNodes function
	linkNodes(previousNode, newNode, nextNode)

	// Check if the previous node's nextPointer is set to the new node
	if previousNode.nextPointer != newNode {
		t.Errorf("Expected previous node's nextPointer to be set to the new node")
	}

	// Check if the new node's previousPointer is set to the previous node
	if newNode.previousPointer != previousNode {
		t.Errorf("Expected new node's previousPointer to be set to the previous node")
	}

	// Check if the new node's nextPointer is set to the next node
	if newNode.nextPointer != nextNode {
		t.Errorf("Expected new node's nextPointer to be set to the next node")
	}

	// Check if the next node's previousPointer is set to the new node
	if nextNode.previousPointer != newNode {
		t.Errorf("Expected next node's previousPointer to be set to the new node")
	}
}
