# Does P Equal NP?

The `P` versus `NP` problem is a major unsolved problem in computer science. It asks whether every problem whose solution can be quickly verified (is in `NP`) can also be solved quickly (is in `P`).

The question is, "Are all `NP` problems really just `P` problems?"

The answer is, "We don't know, but probably not".

## Why do we care?

All problems in `NP` (you know, hard ones like the traveling salesman problem) have been proven to also be solveable in polynomial time if we can find a solution to just one `NP-Complete` problem.

If a single `NP-complete` problem can be solved quickly (in polynomial time) that means that all problems in `NP` can be solved in polynomial time. That would be a huge deal, particularly because it would break digital security systems that rely on the difficulty of certain `NP` problems.