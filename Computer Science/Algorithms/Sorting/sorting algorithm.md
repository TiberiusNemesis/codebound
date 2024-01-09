# Sorting

A sorting algorithm is a sequence of steps to rearrange a sequence of elements into a specific order, usually numerical. The goal of sorting is to allow for easier searching, analysis, and visualization of data. These algorithms are fundamental in computer science, since they enable the execution of other algorithms, such as binary search, linear search, or jump search, which all require sorted data as an input.

Sorting algorithms vary widely in their method of operation, efficiency, memory usage, and stability (whether they preserve the relative order of equal elements). They are generally classified by:

- Time complexity, aka the amount of (computational) time the algorithm takes to complete; some common complexities are $O(n²)$ for simple algorithms like bubble sort, and $O(n log n)$ for more efficient algorithms like quick sort and merge sort.

- Space complexity, or the amount of memory the algorithm requires during execution. Some algorithms (like merge sort) require extra space for temporary storage, while others (like quick sort) sort in place using a minimal amount of additional memory.

- Stability, that is, if it preserves the relative order of records with equal keys (values).

- Adaptivity, in the sense that some algorithms perform better when the input is just _partially_ sorted, making them adaptive.

Examples of sorting algorithms include bubble sort, selection sort, insertion sort, merge sort, quick sort, heap sort, and radix sort. Each has its own strengths and weaknesses.

## Sources

- [GeeksForGeeks, Sorting Algorithms](https://www.geeksforgeeks.org/sorting-algorithms/)
- [freeCodeCamp, Sorting Algorithms Explained](https://www.freecodecamp.org/news/sorting-algorithms-explained-with-examples-in-python-java-and-c/)
