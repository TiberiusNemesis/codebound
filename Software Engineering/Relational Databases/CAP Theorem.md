# CAP Theorem

The CAP Theorem, also known as Brewer's theorem (named after Eric Brewer) says that any system for storing data necessarily can only guarantee two out of these three properties:

- Consistency
- Availability
- Partitioning tolerance

![Venn diagram representing the CAP Theorem](/Assets/cap-theorem.png)

When applied to modern distributed systems, we consider none of these to be 100% safe from network failures, and therefore, generally partition tolerance is taken as a default. We can then only pick consistency or availability as the last factor. 

When choosing consistency over availability, the system will return an error or a time out if particular information cannot be guaranteed to be up to date due to network partitioning. 

And when choosing availability over consistency, the system will always process the query and try to return the most recent available version of the information, even if it cannot guarantee it is up to date due to network partitioning.