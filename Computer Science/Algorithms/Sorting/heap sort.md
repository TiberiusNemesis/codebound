# Heap Sort

An efficient sorting algorithm that uses a binary heap data structure to sort elements. Has better time complexity when compared to bubble sort, selection sort, and insertion sort, especially when dealing with larger sets of data.

It sorts an array by using the properties of a binary heap (i.e. a complete binary tree that satisfies the heap property), most often the max heap.

First, we must convert the input array into a max heap (for sorting in ascending order) or into a min heap (for sorting into a descending order). This is done by heapifying the entire array.

Afterward, we will have the largest (or smallest) element at the root of the heap. We then remove the root element (i.e. swap it with the last element of the heap) and reduce the size of our heap by one. Finally, we heapify the root of the tree again and repeat this process until all elements are sorted.

This algorithm has a time complexity of O(n log n), and since it is an in-place algorithm (i.e. no extra space for data storage besides the original list is required), we can say it has a space complexity of O(1).

## Sources

- [Programiz, Heap Sort](https://www.programiz.com/dsa/heap-sort)
- [GeeksForGeeks, Heap Sort](https://www.geeksforgeeks.org/heap-sort/)
- [Heap sort in 4 minutes](https://www.youtube.com/watch?v=2DmK_H7IdTo)
