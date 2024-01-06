# Quick Sort

A highly-efficient sorting [algorithm](Computer%20Science/Algorithms/algorithm.md) in the average-case performance, this is one of the most commonly used algorithms in practice for larger datasets, although it has a worse-case time complexity when compared to the likes of [bubble sort](Computer%20Science/Algorithms/Sorting/bubble%20sort.md) or [insertion sort](Computer%20Science/Algorithms/Sorting/insertion%20sort.md).

In quick sort, we select an element (called the _pivot_) from the [array](Computer%20Science/Data%20Structures/array.md), then partition the other elements into two other subarrays, according to whether they are greater or lesser than the pivot. Then, these subarrays are sorted recursively -- this can be done in-place, requiring a pretty small amount of memory.

This is an example of a classic _divide and conquer_ [algorithm](Computer%20Science/Algorithms/algorithm.md), which breaks down a single problem into smaller, easier to handle subproblems, then forms the bigger solution from the solved parts.

The best (and average) case scenario is O(n log n). The worst-case is O(n²), and occurs when the pivot is the smallest or the largest element of the list, which leads to very unbalanced partitions. We can _technically_ consider this [algorithm](Computer%20Science/Algorithms/algorithm.md) as having a [logarithmic runtime](Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/logarithmic%20runtime.md), with respect to the input size of the [array](Computer%20Science/Data%20Structures/array.md).

An interesting advantage is that this [algorithm](Computer%20Science/Algorithms/algorithm.md) is highly optimizable; when we apply a strategy for pivot selection and partitioning, we can significantly improve the performance of this already quite efficient sorting method.

## Sources

- [wiki, Quicksort](https://en.wikipedia.org/wiki/Quicksort)
- [Programiz, Quicksort Algorithm](https://www.programiz.com/dsa/quick-sort)
- [GeeksForGeeks, QuickSort](https://www.geeksforgeeks.org/quick-sort/)
