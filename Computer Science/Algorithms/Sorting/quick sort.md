# Quick Sort

A highly-efficient sorting algorithm in the average-case performance, this is one of the most commonly used algorithms in practice for larger datasets, although it has a worse-case time complexity when compared to the likes of bubble sort or insertion sort.

In quick sort, we select an element (called the _pivot_) from the array, then partition the other elements into two other subarrays, according to whether they are greater or lesser than the pivot. Then, these subarrays are sorted recursively -- this can be done in-place, requiring a pretty small amount of memory.

This is an example of a classic _divide and conquer_ algorithm, which breaks down a single problem into smaller, easier to handle subproblems, then forms the bigger solution from the solved parts.

The best (and average) case scenario is O(n log n). The worst-case is O(n²), and occurs when the pivot is the smallest or the largest element of the list, which leads to very unbalanced partitions.

An interesting advantage is that this algorithm is highly optimizable; when we apply a strategy for pivot selection and partitioning, we can significantly improve the performance of this already quite efficient sorting method.

## Sources

- [wiki, Quicksort](https://en.wikipedia.org/wiki/Quicksort)
- [Programiz, Quicksort Algorithm](https://www.programiz.com/dsa/quick-sort)
- [GeeksForGeeks, QuickSort](https://www.geeksforgeeks.org/quick-sort/)
