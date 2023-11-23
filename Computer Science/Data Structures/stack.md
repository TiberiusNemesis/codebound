# Stack
A stack is a type of data structure that describes a linear collection of items -- but operations such as insertion (aka *push*) or deletion (aka *pop*) can only be performed on the last element of the stack. This principle is called "last in, first out", or LIFO.

Stacks are frequently compared to real-life stacks of plates or books, where it is much easier to remove or add a plate (or book) from the top than anywhere else.

In an algorithm, a stack can be used to backtrack to previous steps. In the same vein, many applications use stacks mapping events and actions so that later an user can undo (or redo) the events/actions.

They are very useful whenever:
- elements need to be processed in the reverse order of their insertion (such as function call management, backtracking algorithms, expression evaluation)
- memory management is an important factor to consider
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
- [Stack Data Structure | Illustrated Data Structures](https://www.youtube.com/watch?v=I5lq6sCuABE)
- [Stack in 3 minutes](https://www.youtube.com/watch?v=KcT3aVgrrpU)
- [Stack Data Structure](https://www.coursera.org/lecture/data-structures/stacks-UdKzQ)