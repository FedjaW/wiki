# The two crystal balls problem

Problem:

Given two crystal balls that will break if dropped from high enough distance, determine the exact spot in which it will break in the most optimized way?

Reformulated in as an array:

[ 0 0 0 0 1 1 1 1 1 1 1 ]
0 . . . . . . . . . . . . . . N

If we pick the mid and the first ball breaks, we have to start from 0 and go one by one until the second ball breaks. Then we have the solution, but we have O(1/2N) which is O(N) which is the same as if we would use linear search and had only one ball. So no win here.

Answere:

We jump by Sqrt(N). The worst case is to jump Sqrt(N) and then go back to the last known point and walk Sqrt(N).

That is:

O(Sqrt(N) + Sqrt(N)) = O(2*Sqrt(N)) = O(Sqrt(N))

```j
// breaks array is has this format: [0,0,1,1,1,1,1]

// my own solution
export default function two_crystal_balls(breaks: boolean[]): number {
    const jumpAmount = Math.floor(Math.sqrt(breaks.length));
    const numberOfJumps = breaks.length / jumpAmount

    var lowerEnd = 0;

    for (let n = 1; n < numberOfJumps + 1; ++n) {
        const index = jumpAmount * n

        if (breaks[index] === true) {
            for (let i = lowerEnd; i < index; i++) {
                if (breaks[i] === true) {
                    return i;
                }
                lowerEnd = i;
            }
        }
    }

    return -1 // indicating no match aka crystal ball won't break at all
}

// primeagen's solution
function two_crystal_balls_primeagen(breaks: boolean[]): number {
    const jmpAmount = Math.floor(Math.sqrt(breaks.length));
    let i = jmpAmount

    for (; i < breaks.length; i += jmpAmount) {
        if (breaks[i]) {
            break
        }
    }

    i -= jmpAmount

    for (let j = 0; j < jmpAmount && i < breaks.length; ++j, ++i) {
        if (breaks[i]) {
            return i
        }
    }

    return -1 // indicating no match aka crystal ball won't break at all
}

```