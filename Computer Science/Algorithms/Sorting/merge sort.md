# Merge Sort

A sorting [algorithm](Computer%20Science/Algorithms/algorithm.md) which continuously divides elements into arrays and subarrays, then combines (or _merges_) them back together into a single, sorted [array](Computer%20Science/Data%20Structures/array.md). This is one form of a _divide and conquer_ [algorithm](Computer%20Science/Algorithms/algorithm.md), which breaks down a single problem into smaller, easier to handle subissues, then forms the bigger solution from the solved subproblems.

To properly [merge](Software%20Engineering/Version%20Control%20Systems/merge.md) sort, we must divide the [array](Computer%20Science/Data%20Structures/array.md) into halves further and further until we have a base case of an [array](Computer%20Science/Data%20Structures/array.md) with only one element. Then, we begin merging the arrays in a specific order, making sure that whatever arrays we have are sorted. So when merging arrays of any size, we know the result will be sorted; by the end of the sorting process, we will have a single, fully sorted [array](Computer%20Science/Data%20Structures/array.md).

Merge sort has a [logarithmic runtime](Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/logarithmic%20runtime.md), with a worst-case scenario of _log(n * log n)_.

## Sources

- [GeeksForGeeks, Merge Sort](https://www.geeksforgeeks.org/merge-sort/)
- [Programiz, Merge Sort Algorithm](https://www.programiz.com/dsa/merge-sort)
