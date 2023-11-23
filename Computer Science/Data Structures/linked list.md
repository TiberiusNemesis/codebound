# Linked List
A [data structure](Computer%20Science/Data%20Structures/data%20structure.md) which is made up of **nodes.**  
Each node consists of two elements: the data (or value), and a pointer to the address of another node in the list. 

This relationship allows us to form a list of elements where, unlike an [array](Computer%20Science/Data%20Structures/array.md), the elements do not need to be sequential in the computer's memory.

They are very useful whenever:
- the total number of elements is unknown in advance
- insertion or deletion operations are more frequent than element access
- memory efficiency is important (since each element is allocated when the space is needed, not before the necessity arises)
## How does it work?
We create (and access) the linked list through a reference node; this node is either in the start of the list or at the end of it. We call these the *head* and *tail* nodes of the list. 

This node will contain some data (e.g. a string, a number, an object...) and at least one pointer that leads to another element. The next element will also contain data and another pointer. 

When reaching the tail of the list, that node will then not contain a memory address; instead, it will point to null.

If each node points only to the next element, we call this a singly linked list.
Otherwise, if each node points to both the previous and the next element, we call it a doubly linked list.
## Operations and Time Complexity
Inserting an element into a linked list is at most O(1) complexity because the elements are not stored sequentially in memory, which means we can allocate memory to create a node and just change the pointers from two elements in the list. The same goes for deleting a node -- all that's needed is to remap the pointers from the nodes that directly precede and succeed it.

Transversing a linked list and accessing one of its elements, however, is (at worst) O(n) complexity since when starting from the head (or tail) of the list, it is necessary to follow the pointer references to each next element until we get to the ones we are looking to displace.

*The summary below refers to doubly linked lists.*
### Worst-case 
- Access: O(n)
- Search: O(n)
- Insertion: O(1)
- Deletion: O(1)
### Best-case
- Access: O(1)
- Search: O(1)
- Insertion: O(1)
- Deletion: O(1)
## Sources
- [Linked List Data Structure | Illustrated Data Structures](https://www.youtube.com/watch?v=odW9FU8jPRQ)
- [Linked Lists in 4 minutes](https://www.youtube.com/watch?v=F8AbOfQwl1c)
- [Singly Linked Lists](https://www.coursera.org/lecture/data-structures/singly-linked-lists-kHhgK)
- [Linked Lists (USP)](https://www.ime.usp.br/~pf/algorithms/chapters/linked-lists.html#:~:text=Linked%20lists,second%20cell%2C%20and%20so%20on.)
- [wiki: Linked list](https://en.wikipedia.org/wiki/Linked_list)