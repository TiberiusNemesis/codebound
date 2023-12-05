# Big O Notation

Big O is a type of asymptotic notation we use to refer to the runtime of an [algorithm](Computer%20Science/Algorithms/algorithm.md) in its worst-case scenario.

For example:

You have an [algorithm](Computer%20Science/Algorithms/algorithm.md) that sorts arrays and you input an [array](Computer%20Science/Data%20Structures/array.md) with 5 elements, and the [algorithm](Computer%20Science/Algorithms/algorithm.md) takes 5 s steps (or passes) to sort it.

Then you input 100 elements, and the [algorithm](Computer%20Science/Algorithms/algorithm.md) takes 100 steps (or passes) to sort it.  
You can say this [algorithm](Computer%20Science/Algorithms/algorithm.md) is _linear,_ i.e. the resources (usually time taken) scale linearly with the amount of elements of the input.
In Big O, we'd represent this as O(n).

Other common runtimes include:

- constant, aka constant number of steps e.g. O(1), O(5)
- logarithmic, O(log n)
- polynomial, O(nᴷ), e.g. quadratic, O(n²)
- exponential, O(2ⁿ)
- factorial, O(n!)