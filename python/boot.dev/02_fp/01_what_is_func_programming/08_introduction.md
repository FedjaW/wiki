# What is Functional Programming?

Functional programming is a style (or "paradigm" if you're pretentious) of programming where we compose functions instead of mutating state (updating the value of variables).

- Functional programming is more about declaring what you want to happen, rather than how you want it to happen.
- Imperative (or procedural) programming declares both the what and the how.

Example of imperative code:

```python
car = create_car()
car.add_gas(10)
car.clean_windows()
```

Example of functional code:

```python
return clean_windows(add_gas(create_car()))
```

The important distinction is that in the functional example, we never change the value of the car variable, we just compose functions that return new values, with the outermost function, clean_windows in this case, returning the final result.

## Why Python?

Frankly, Python is not the best language for functional programming. Reasons include:

1. No static typing.
2. Everything is mutable.
3. No tail call optimization.
4. Side effects are common.
5. Imperative and OOP styles abound in popular libraries.
6. Purity is not enforced (and sometimes not even encouraged).
7. Sum Types are hard to define.
8. Pattern matching is weak at best.

So seriously, why are we using Python? One reason trumps all others: you already know Python. Python is a great choice for learning coding basics, OOP, Algorithms, and Data Structures, and the tradeoff of learning a new language at this point in the curriculum isn't worth it.

We can still cover the most important concepts of functional programming in Python, even if we have to jump through a hoop or two to do it. Functional programming is a paradigm of useful techniques for writing better code, and they apply to all languages, not just purely functional ones.

Note: We also plan to release a "Functional Programming 2" course in a more functional language. Likely one of these:

- Haskell
- OCaml
- Elixir