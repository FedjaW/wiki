# Sizeof

C gives us a way to check the size of a type or a variable: sizeof.

You can use sizeof like a function (although, technically it's a unary operator, but that distinction is generally only useful for winning super important internet arguments).

We'll use the sizeof operator in the next few lessons to give us insight into the memory layout of different types in C. This will be particularly useful as we move deeper into C, and essential for understanding pointers.

> Pointers are not too bad once you understand the basics!

## size_t

The size_t type is a special type that is guaranteed to be able to represent the size of the largest possible object in the target platform's address space (i.e. can fit any single, non-struct value inside of it).

It's also the type that sizeof returns.