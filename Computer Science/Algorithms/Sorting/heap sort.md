# Heap Sort

An efficient sorting [algorithm](Computer%20Science/Algorithms/algorithm.md) that uses a binary [heap](Computer%20Science/Data%20Structures/heap.md) [data structure](Computer%20Science/Data%20Structures/data%20structure.md) to sort elements. Has better time complexity when compared to [bubble sort](Computer%20Science/Algorithms/Sorting/bubble%20sort.md), [selection sort](Computer%20Science/Algorithms/Sorting/selection%20sort.md), and [insertion sort](Computer%20Science/Algorithms/Sorting/insertion%20sort.md), especially when dealing with larger sets of data.

It sorts an [array](Computer%20Science/Data%20Structures/array.md) by using the properties of a binary [heap](Computer%20Science/Data%20Structures/heap.md) (i.e. a [complete binary tree](Computer%20Science/Data%20Structures/Tree/complete%20binary%20tree.md) that satisfies the [heap](Computer%20Science/Data%20Structures/heap.md) property), most often the max [heap](Computer%20Science/Data%20Structures/heap.md).

First, we must convert the input [array](Computer%20Science/Data%20Structures/array.md) into a max [heap](Computer%20Science/Data%20Structures/heap.md) (for sorting in ascending order) or into a min [heap](Computer%20Science/Data%20Structures/heap.md) (for sorting into a descending order). This is done by heapifying the entire [array](Computer%20Science/Data%20Structures/array.md).

Afterward, we will have the largest (or smallest) element at the root of the [heap](Computer%20Science/Data%20Structures/heap.md). We then remove the root element (i.e. swap it with the last element of the [heap](Computer%20Science/Data%20Structures/heap.md)) and reduce the size of our [heap](Computer%20Science/Data%20Structures/heap.md) by one. Finally, we heapify the root of the [tree](Computer%20Science/Data%20Structures/tree.md) again and repeat this process until all elements are sorted.

This [algorithm](Computer%20Science/Algorithms/algorithm.md) has a time complexity of O(n log n), and since it is an in-place [algorithm](Computer%20Science/Algorithms/algorithm.md) (i.e. no extra space for data storage besides the original list is required), we can say it has a space complexity of O(1).

## Sources

- [Programiz, Heap Sort](https://www.programiz.com/dsa/heap-sort)
- [GeeksForGeeks, Heap Sort](https://www.geeksforgeeks.org/heap-sort/)
- [Heap sort in 4 minutes](https://www.youtube.com/watch?v=2DmK_H7IdTo)
