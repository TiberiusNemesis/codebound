# Selection Sort

A basic in-place sorting [algorithm](Computer%20Science/Algorithms/algorithm.md) where we divide the [array](Computer%20Science/Data%20Structures/array.md) into sorted and unsorted -- in the unsorted portion, we pass through the list and find the smallest (or largest, depending on what we're sorting by) element, and swap it with the leftmost element from the unsorted list, thus adding +1 element to the sorted portion.

This [algorithm](Computer%20Science/Algorithms/algorithm.md) has a time complexity of O(n²), representing a [polynomial runtime](Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/polynomial%20runtime.md), due to the fact only one element is guaranteed to be sorted per pass.

## Sources

- [Selection sort in 3 minutes](https://www.youtube.com/watch?v=g-PGLbMth_g)
- [wiki, Selection sort](https://en.wikipedia.org/wiki/Selection_sort)
- [GeeksForGeeks, Selection Sort](https://www.geeksforgeeks.org/selection-sort/)
