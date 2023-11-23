# Queue
A [data structure](Computer%20Science/Data%20Structures/data%20structure.md) where the elements stored are inserted and deleted in a particular order, following the first in, first out (or FIFO) principle, *i.e.* the first element to be inserted is the first element to be removed. 

They are very useful whenever:
- a strict order of processing elements is required (such as task scheduling, breadth-first search, handling requests in a web server)
- concurrency/async/multithreading is a factor (since queues can help in synchronizing and coordinating tasks between threads and processes)
## Operations and Time Complexity

### Worst-case
- Access: O(n)
- Search: O(n)
- Insertion: O(1)
- Deletion: O(1)
### Best-case
- Access: O(1)
- Search: O(1)
- Insertion: O(1)
- Deletion: O(1)
## Sources
- [Queue Data Structure | Illustrated Data Structures](https://www.youtube.com/watch?v=mDCi1lXd9hc)
- [Queue in 3 Minutes](https://www.youtube.com/watch?v=D6gu-_tmEpQ)
- [Circular Buffer - Wikipedia](https://en.wikipedia.org/wiki/Circular_buffer)