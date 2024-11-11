# NP-Complete

Some, but not all problems in `NP` are also NP-complete.

A problem in `NP` is also `NP-complete` if every other problem in `NP` can be reduced to it in polynomial time.

## What does "reduced" mean?

We won't dive deep into the subject of reductions in this course, but we'll cover the basic idea.

Any problem, let's call it Problem A, can be reduced to a different problem, Problem B, if there is an algorithm (called a reducer) that transforms an algorithm that solves Problem B into an algorithm that solves Problem A.

Algorithm to solve B -> reducer -> Algorithm to solve A

However, the reducer itself needs to be fast. "Problem A is reducible to Problem B" if the reducer can run in polynomial time.

## So who cares?

Well, this means that if we can find an algorithm that solves any of the NP-complete problems in polynomial time, then all problems in NP can also be solved in polynomial time.

Super-duper-smart computer scientists have proven it. Trust me. Or optionally read more about it if you're interested: <https://web.stanford.edu/class/archive/cs/cs103/cs103.1134/lectures/26/Small26.pdf>
