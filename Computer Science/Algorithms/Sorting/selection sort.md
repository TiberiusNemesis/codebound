# Selection Sort

A basic in-place sorting algorithm where we divide the list into sorted and unsorted -- in the unsorted portion, we pass through the list and find the smallest (or largest, depending on what we're sorting by) element, and swap it with the leftmost element from the unsorted list, thus adding +1 element to the sorted portion.

This algorithm has a time complexity of O(n²), representing a polynomial runtime, due to the fact only one element is guaranteed to be sorted per pass.

## Sources

- [Selection sort in 3 minutes](https://www.youtube.com/watch?v=g-PGLbMth_g)
- [wiki, Selection sort](https://en.wikipedia.org/wiki/Selection_sort)
- [GeeksForGeeks, Selection Sort](https://www.geeksforgeeks.org/selection-sort/)
