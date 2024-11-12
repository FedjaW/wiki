# Why Pointers?

To illustrate the usefulness of pointers, let's pretend we want to pass a collection of data into a function. Within that function, we want to modify the data. In Python, we could use a class to store the data, and pass an instance of that class into the function:

```py
class Coordinate:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

def update_coordinate_x(coord, new_x):
    coord.x = new_x

c = Coordinate(1, 2, 3)
print(c.x)  # 1
update_coordinate_x(c, 4)
print(c.x)  # 4
```

## Assignment

Now let's do the same thing, but using a struct in C.

Complete the `coordinate_update_x` and `coordinate_update_and_return_x` functions.

`coordinate_update_x` should update the `x` field with the provided `new_x` value. It returns void.
`coordinate_update_and_return_x` should update the `x` field with the provided new_x value, and then return the updated coordinate struct.
Remember, you can access the field of a struct with a `.` operator, like so:

```c
car.tires = 4;
```

### Code

```c
#include "coordinate.h"

void coordinate_update_x(coordinate_t coord, int new_x) {
  coord.x = new_x;
}

coordinate_t coordinate_update_and_return_x(coordinate_t coord, int new_x) {
  coord.x = new_x;
  return coord;
}
```

### What happened?

After passing the assignment, open up `main.c` and take a look at the test cases. You'll notice that `coordinate_update_x` doesn't update anything, but `coordinate_update_and_return_x` does because it returns a new copy of the struct.

In C, structs are passed by value. That's why updating a field in the struct does not change the original struct from the `main` function.
To get the change to "persist", we needed to return the updated struct from the function (a new copy).
The memory address of the struct that went in to `coordinate_update_and_return_x` was not the same as the address of the struct that was returned. Again, because we created a copy.

#### Main.c

```c
#include "munit.h"

#include "coordinate.h"

coordinate_t new_coordinate(int x, int y, int z) {
  return (coordinate_t){.x = x, .y = y, .z = z};
}

munit_case(RUN, test_unchanged, {
  coordinate_t old = new_coordinate(1, 2, 3);
  munit_assert_int(old.x, ==, 1, "old.x must be 1");

  coordinate_update_x(old, 4);
  munit_assert_int(old.x, ==, 1, "old.x must still be 1");
});

munit_case(SUBMIT, test_changed, {
  coordinate_t old = new_coordinate(1, 2, 3);
  munit_assert_int(old.x, ==, 1, ".x must be 1");

  coordinate_t new = coordinate_update_and_return_x(old, 4);
  munit_assert_int(new.x, ==, 4, "new .x must be 4");
  munit_assert_int(old.x, ==, 1, "old.x must still be 1");

  // Notice, they have different addresses
  munit_assert_ptr_not_equal(&old, &new, "Must be different addresses");
});

int main() {
  MunitTest tests[] = {
      munit_test("/test_unchanged", test_unchanged),
      munit_test("/test_changed", test_changed),
      munit_null_test,
  };

  MunitSuite suite = munit_suite("pointers", tests);

  return munit_suite_main(&suite, NULL, 0, NULL);
}
```