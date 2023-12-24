# Big O Notation

Big O is a type of asymptotic notation we use to refer to the complexity of an [algorithm](Computer%20Science/Algorithms/algorithm.md) in its worst-case scenario. This is opposed to [Big Ω](Computer%20Science/Asymptotic%20Notation/Big%20Ω.md), which represents the best-case scenario, and is complemented by [Big Θ](Computer%20Science/Asymptotic%20Notation/Big%20Θ.md), which represents the average-case scenario.

For example:

You have an [algorithm](Computer%20Science/Algorithms/algorithm.md) that sorts arrays and you input an [array](Computer%20Science/Data%20Structures/array.md) with 5 elements, and the [algorithm](Computer%20Science/Algorithms/algorithm.md) takes 5 s steps (or passes) to sort it.

Then you input 100 elements, and the [algorithm](Computer%20Science/Algorithms/algorithm.md) takes 100 steps (or passes) to sort it.  
You can say this [algorithm](Computer%20Science/Algorithms/algorithm.md) is _[linear](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/linear%20runtime.md),_ i.e. the resources (usually time taken) scale linearly with the amount of elements of the input.
In Big O, we'd represent this as O(n).

Other common runtimes include:

- [constant](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/constant%20runtime.md), aka constant number of steps e.g. O(1), O(5)
- [logarithmic](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/logarithmic%20runtime.md), O(log n)
- [polynomial](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/polynomial%20runtime.md), O(nᴷ), e.g. quadratic, O(n²)
- [exponential](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/exponential%20runtime.md), O(2ⁿ)
- [factorial](/Computer%20Science/Asymptotic%20Notation/Common%20Runtimes/factorial%20runtime.md), O(n!)

Here's a visual representation.

![Image showing a graph of algorithmic complexity runtimes ](/Assets/runtime-complexity.png)

## Sources

- [Big O Notation — Calculating Time Complexity](https://www.youtube.com/watch?v=Z0bH0cMY0E8)
- [Big O Notations](https://www.youtube.com/watch?v=V6mKVRU1evU)
