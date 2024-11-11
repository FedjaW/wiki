# Strings

Most programming languages these days (even compiled ones) have a built-in `string` type of some sort. C... doesn't.

Instead, C strings are just arrays (like lists) of characters. We'll talk more about the specifics when we talk about arrays and pointers later, but for now know that this is how you get a "string" in C:

```c
char *msg_from_dax = "You still have 0 users";
```

Very (I repeat, very) loosely speaking, `char *` means string. Also note that it is required to use double quotes `"`. Single quotes (`'`) make char, not `char *`.