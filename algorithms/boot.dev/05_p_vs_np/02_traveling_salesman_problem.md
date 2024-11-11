# Traveling Salesman Problem

A famous example of a problem in NP is the Traveling Salesman Problem, also known as TSP.

The version of the problem that we will solve can be stated:

> Given a list of cities, the distances between each pair of cities, and a total distance, is there a path through all the cities that is less than the distance given?

## Pseudocode

Inputs:

- cities: A list of city identifiers (integers 0-n)
- paths: A matrix where each point represents the distance between the two cities. For example, paths[cityA][cityB] holds the distance from cityA to cityB. paths[cityA][cityB] = paths[cityB][cityA]
- dist: The distance we are trying to beat

## Algorithm

1. Use the permutations function to get the matrix of all possible paths through the given cities. Where the first path, [1, 2, 3] represents moving from city 1 -> city 2 -> city 3
2. Iterate over each possible path (permutation)
    1. Sum the distances between each city in the path using the paths matrix to look up the distances
    2. If the total distance of the path is less than the given dist return True
3. If no short paths were found, return False

## Implementation

```py
def tsp(cities, paths, dist):
    perms = permutations(cities)
    for perm in perms:
        total_dist = 0
        for i in range(1, len(perm)):
            total_dist += paths[perm[i - 1]][perm[i]]
        if total_dist < dist:
            return True
    return False

def permutations(arr):
    res = []
    res = helper(res, arr, len(arr))
    return res


def helper(res, arr, n):
    if n == 1:
        tmp = arr.copy()
        res.append(tmp)
    else:
        for i in range(n):
            res = helper(res, arr, n - 1)
            if n % 2 == 1:
                arr[n - 1], arr[i] = arr[i], arr[n - 1]
            else:
                arr[0], arr[n - 1] = arr[n - 1], arr[0]
    return res
```

## TSP in NP - Verification

Although it takes factorial time to solve `TSP`, `TSP` is in `NP` because we can verify a solution much faster. Let's write a `TSP` verifier!

### Pseudocode - Verification

Inputs:

- paths: A matrix where each point represents the distance between the two cities. For example, paths[cityA][cityB] holds the distance from cityA to cityB. paths[cityA][cityB] = paths[cityB][cityA]
- dist: The distance we are trying to find a path shorter than
- actual_path: The path we are trying to verify

### Algorithm - Verification

1. Loop over each city in the actual path
2. Sum the distance between each city in the actual path
3. If the sum is less than dist, return True, otherwise return False

### Implementation - Verification

```py
def verify_tsp(paths, dist, actual_path):
    total = 0
    for i in range(len(actual_path)):
        if i != 0:
            total += paths[actual_path[i-1]][actual_path[i]]
    return total < dist
```

The Big O of 'tsp' is `O(n!)` and 'verify_tsp' is `O(n)`.
