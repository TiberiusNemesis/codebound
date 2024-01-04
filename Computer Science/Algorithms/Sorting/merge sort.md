# Merge Sort

A sorting algorithm which continuously divides elements into arrays and subarrays, then combines (or _merges_) them back together into a single, sorted array. This is one form of a _divide and conquer_ algorithm, which breaks down a single problem into smaller, easier to handle subissues, then forms the bigger solution from the solved subproblems.

To properly merge sort, we must divide the array into halves further and further until we have a base case of an array with only one element. Then, we begin merging the arrays in a specific order, making sure that whatever arrays we have are sorted. So when merging arrays of any size, we know the result will be sorted; by the end of the sorting process, we will have a single, fully sorted array.

Merge sort has a linear runtime, with a worst-case scenario of _log(n * log n)_.

## Sources

- [GeeksForGeeks, Merge Sort](https://www.geeksforgeeks.org/merge-sort/)
- [Programiz, Merge Sort Algorithm](https://www.programiz.com/dsa/merge-sort)
