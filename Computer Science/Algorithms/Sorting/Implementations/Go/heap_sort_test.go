package sorting

import (
	"reflect"
	"testing"

	ds "github.com/TiberiusNemesis/codebound/ComputerScience/DataStructures/Implementations/Go"
)

func isMaxHeap(node *ds.BinaryTreeNode) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && node.Left.Key > node.Key {
		return false
	}

	if node.Right != nil && node.Right.Key > node.Key {
		return false
	}

	return isMaxHeap(node.Left) && isMaxHeap(node.Right)
}

func TestBuildMaxHeap(t *testing.T) {
	tests := []struct {
		elements []int
	}{
		{[]int{2, 5, 55, 10, 24, 33, 1}},
		{[]int{10, 5, 2, 1, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{30, 20, 10}},
	}

	for _, test := range tests {
		tree := ds.BuildCompleteBinaryTree(test.elements)
		BuildMaxHeap(tree)

		if !isMaxHeap(tree) {
			t.Errorf("BuildMaxHeap failed to create a max heap from the tree %+v", test.elements)
		}
	}
}

func TestDeleteRoot(t *testing.T) {
	tests := []struct {
		elements []int
		isValid bool
	}{
		{[]int{55, 24, 33, 10, 5, 2, 1}, true},
		{[]int{5, 4, 3, 2, 1}, true},
		{[]int{1}, true},
		{[]int{}, true},
	}

	for _, test := range tests {
		tree := ds.BuildCompleteBinaryTree(test.elements)
		btree := &ds.BinaryTree{Root: tree}
		BuildMaxHeap(tree)
		DeleteRoot(btree)
		isValidHeap := isMaxHeap(btree.Root)

		if (isValidHeap != test.isValid) {
			t.Errorf("DeleteRoot failed for tree %+v, expected %+v, got %+v", test.elements, test.isValid, isValidHeap)
		}
	}
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		elements []int
		expected []int // Expected sorted array
	}{
		{[]int{55, 24, 33, 10, 5, 2, 1}, []int{1, 2, 5, 10, 24, 33, 55}},
		{[]int{3, 1, 4, 1, 5, 9, 2}, []int{1, 1, 2, 3, 4, 5, 9}},
		{[]int{1}, []int{1}}, // Single element
		{[]int{}, []int{}},   // Empty tree
	}

	for _, test := range tests {
		tree := ds.BuildCompleteBinaryTree(test.elements)
		sortedArr := HeapSort(&ds.BinaryTree{Root: tree})

		if !reflect.DeepEqual(sortedArr, test.expected) {
			t.Errorf("HeapSort failed for tree %+v, expected %+v, got %+v", test.elements, test.expected, sortedArr)
		}
	}
}
