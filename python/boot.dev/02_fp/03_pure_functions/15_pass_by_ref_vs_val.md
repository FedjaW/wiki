# Reference vs. Value

When you pass a value into a function as an argument, one of two things can happen:

It's passed by reference: The function has access to the original value and can change it
It's passed by value: The function only has access to a copy. Changes to the copy within the function don't affect the original
There is a bit more nuance, but this explanation mostly works.

These types are passed by reference:

- Lists
- Dictionaries
- Sets

These types are passed by value:

- Integers
- Floats
- Strings
- Booleans
- Tuples

Most collection types are passed by reference (except for tuples) and most primitive types are passed by value.
