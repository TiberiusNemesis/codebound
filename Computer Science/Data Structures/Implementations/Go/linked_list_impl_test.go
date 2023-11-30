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
	// TODO: Write test cases for the getPreviousNode method
}

func TestListNode_GetNextNode(t *testing.T) {
	// TODO: Write test cases for the getNextNode method
}

func TestLinkNodes(t *testing.T) {
	// TODO: Write test cases for the linkNodes function
}
