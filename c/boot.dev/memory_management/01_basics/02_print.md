# Printing Variables

You've already seen me do a `printf()` magic a few times. Unfortunately, in C it isn't as easy to do string interpolation (what f-strings do in Python).

Instead of:

```py
print(f"Hello, {name}. You're {age} years old.")
```

We have to tell C how we want particular values to be printed using "format specifiers".

Common format specifiers are:

- %d - digit (integer)
- %c - character
- %f - floating point number
- %s - string (char *)

```c
printf("Hello, %s. You're %d years old.\n", name, age);
```

## Newline Character

The `print()` function in Python automatically adds a newline character (\n) at the end of the string. In C, we have to do this manually.

```c
printf("Hello, world!\n");
```

