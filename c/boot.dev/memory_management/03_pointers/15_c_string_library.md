# C String Library

The C standard library provides a comprehensive set of functions to manipulate strings in the `<string.h>` header file. Here are some of the most commonly used functions:

- `strcpy`: Copies a string to another.

```c
char src[] = "Hello";
char dest[6];
strcpy(dest, src);
// dest now contains "Hello"
```

- `strcat`: Concatenates (appends) one string to another.

```c
char dest[12] = "Hello";
char src[] = " World";
strcat(dest, src);
// dest now contains "Hello World"
```

- `strlen`: Returns the length of a string (excluding the null terminator).

```c
char str[] = "Hello";
size_t len = strlen(str);
// len is 5
```

- `strcmp`: Compares two strings lexicographically.

```c
char str1[] = "Hello";
char str2[] = "World";
int result = strcmp(str1, str2);
// result is negative since "Hello" < "World"
```

- `strncpy`: Copies a specified number of characters from one string to another.

```c
char src[] = "Hello";
char dest[6];
strncpy(dest, src, 3);
// dest now contains "Hel"
dest[3] = '\0';
// ensure null termination
```

- `strncat`: Concatenates a specified number of characters from one string to another.

```c
char dest[12] = "Hello";
char src[] = " World";
strncat(dest, src, 3);
// dest now contains "Hello Wo"
```

- `strchr`: Finds the first occurrence of a character in a string.

```c
char str[] = "Hello";
char *pos = strchr(str, 'l');
// pos points to the first 'l' in "Hello"
```

- `strstr`: Finds the first occurrence of a substring in a string.

```c
char str[] = "Hello World";
char *pos = strstr(str, "World");
// pos points to "World" in "Hello World"
```