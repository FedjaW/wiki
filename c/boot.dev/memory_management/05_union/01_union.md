# Union

Now that we understand structs and enums, we can learn aboutunions: a combination of the two concepts.

Unions in C can hold one of several types. They're like a less-strict sum type from the world of functional programming. Here's an example union:

```c
typedef union AgeOrName {
  int age;
  char *name;
} age_or_name_t;
```

The age_or_name_t type can hold either an int or a char *, but not both at the same time (that would be a struct). We provide the list of possible types so that the C compiler knows the maximum potential memory requirement, and can account for that. This is how the union is used:

```c
age_or_name_t lane = { .age = 29 };
printf("age: %d\n", lane.age);
// age: 29
```

Here's where it gets interesting. What happens if we try to access the name field (even though we set the age field)?

```c
printf("name: %s\n", lane.name);
// name:
```

We get... nothing? To be more specific, we get undefined behavior. A union only reserves enough space to hold the largest type in the union and then all of the fields use the same memory. So when we set .age to 29, we are writing the integer representation of 29 to the memory of the lane union:

```txt
0000 0000 0000 0000 0000 0000 0001 1101
```

Then if we try to access .name, we read from the same block of memory but try to interpret the bytes as a char *, which is why we get garbage (which is interpreted as nothing in this case). Put simply, setting the value of .age overwrites the value of .name and vice versa, and you should only access the field that you set.

## Memory Layout

Unions store their value in the same memory location, no matter which field or type is actively being used. That means that accessing any field apart from the one you set is generally a bad idea.
