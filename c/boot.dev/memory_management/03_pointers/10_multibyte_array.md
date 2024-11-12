# Multibyte Arrays

If we create an array of structs it gets crazy because we can access and manipulate the elements using either indexing or pointer arithmetic. Let's see how multi-byte width structures are managed in memory.

First, let's say we're working with our familiar Coordinate struct:

```c
typedef struct Coordinate {
  int x;
  int y;
  int z;
} coordinate_t;
```

We can declare an array of 3 Coordinate structs like so:

```c
coordinate_t points[3] = {
  {1, 2, 3},
  {4, 5, 6},
  {7, 8, 9}
};
```

Then we can print out the values of the second element in the array:

```c
printf("points[1].x = %d, points[1].y = %d, points[1].z = %d\n",
  points[1].x, points[1].y, points[1].z
);
// 4, 5, 6
```

Or we can use a pointer:

```c
coordinate_t *ptr = points;
printf("ptr[1].x = %d, ptr[1].y = %d, ptr[1].z = %d\n",
  (ptr + 1)->x, (ptr + 1)->y, (ptr + 1)->z
);
// 4, 5, 6
```