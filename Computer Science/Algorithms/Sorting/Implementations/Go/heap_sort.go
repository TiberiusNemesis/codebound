package sorting

import (
    "fmt"
    ds "github.com/TiberiusNemesis/codebound/ComputerScience/DataStructures/Implementations/Go"
)

//////// Using arrays

// buildMaxHeap creates a max heap from an unsorted array
func buildMaxHeap (arr []int) {
    n := len(arr)
    for i := (n/2)-1; i >= 0; i-- {
        heapify(arr, n, i)
    }
}

// heapify is similar to build max heap, but assumes part of the array is already sorted
func heapify (arr []int, n int, i int) {
    left := 2 * i + 1
    right := 2 * i + 2
    largest := i

    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

// heapSort sorts an array in ascending order (uses max heap)
func heapSort(arr []int) {
    n := len(arr)

    buildMaxHeap(arr)

    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
}

//////// Using binary tree 

// Heapify recursively maintains the max-heap property starting from a given node
func Heapify(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }

    largest := node
    left := node.Left
    right := node.Right

    if left != nil && left.Key > largest.Key {
        largest = left
    }

    if right != nil && right.Key > largest.Key {
        largest = right
    }

    if largest != node {
        node.Key, largest.Key = largest.Key, node.Key
        Heapify(largest)
    }
}

// heapifyDown restructures the heap from a node down to maintain the max-heap property.
func heapifyDown(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }

    largest := node
    if node.Left != nil && node.Left.Key > largest.Key {
        largest = node.Left
    }
    if node.Right != nil && node.Right.Key > largest.Key {
        largest = node.Right
    }

    if largest != node {
        node.Key, largest.Key = largest.Key, node.Key
        heapifyDown(largest)
    }
}

// BuildMaxHeap builds a max-heap recursively using a node (ideally starting from a root node)
func BuildMaxHeap(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }
    // Perform a post-order traversal starting from the leaves and moving up.
    // This ensures that the subtree under each node is already a max-heap.
    BuildMaxHeap(node.Left)
    BuildMaxHeap(node.Right)

    // Heapify the current node to maintain the max-heap property.
    Heapify(node)
}

// HeapSort sorts the binary tree using heap sort and returns a sorted array.
func HeapSort(tree *ds.BinaryTree) []int {
    if tree == nil || tree.Root == nil {
        return []int{}
    }

    size := tree.Size()
    sortedArr := make([]int, size)

    BuildMaxHeap(tree.Root)

    for i := 0; i < size; i++ {
        sortedArr[size - 1 - i] = tree.Root.Key
        DeleteRoot(tree)
    }

    return sortedArr
}


// DeleteRoot removes the root node of the binary heap and maintains the heap property.
func DeleteRoot(t *ds.BinaryTree) {
    if t.Root == nil {
        return
    }

    // Find the last node in the binary heap.
    lastNode := ds.GetLastNode(t.Root, t.Size())
    
    // Replace the root's value with the last node's value.
    t.Root.Key = lastNode.Key

    // Remove the last node.
    t.Root = ds.RemoveLastNode(t.Root, t.Size())

    // Heapify down from the root to maintain the heap property.
    heapifyDown(t.Root)
}

// InorderTraversal performs an inorder traversal of the binary tree for testing.
func InorderTraversal(node *ds.BinaryTreeNode) {
    if node == nil {
        return
    }
    InorderTraversal(node.Left)
    fmt.Printf("%d ", node.Key)
    InorderTraversal(node.Right)
}