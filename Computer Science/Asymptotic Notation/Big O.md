# Big O Notation

Big O is a type of asymptotic notation we use to refer to the runtime of an algorithm in its worst-case scenario.

For example:

You have an algorithm that sorts arrays and you input an array with 5 elements, and the algorithm takes 5 s steps (or passes) to sort it.

Then you input 100 elements, and the algorithm takes 100 steps (or passes) to sort it.  
You can say this algorithm is _linear,_ i.e. the resources (usually time taken) scale linearly with the amount of elements of the input.
In Big O, we'd represent this as O(n).

Other common runtimes include:

- constant, aka constant number of steps e.g. O(1), O(5)
- logarithmic, O(log n)
- polynomial, O(nᴷ), e.g. quadratic, O(n²)
- exponential, O(2ⁿ)
- factorial, O(n!)
