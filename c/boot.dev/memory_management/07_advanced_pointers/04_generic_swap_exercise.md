# Generic Swap

To make a generic `swap`, we will need to provide the C compiler with the size of the data that we are swapping because `void *` loses that type info. Our new interface for `swap` will include the size:

```c
void swap(void *vp1, void *vp2, size_t size);
```

The other problem we're going to have is that directly assigning pointer values does not work the same way with `void *`. Instead of using `*ptr1 = *ptr2`, we will use `memcpy`, which is included in the string.h library:

```c
void *memcpy(void *destination, void* source, size_t size);
```

So to move the data from `ptr2` to `ptr1`, we will use the following:

```c
memcpy(ptr1, ptr2, size);
```

## Implementation

```c
#include <stdlib.h>
#include <string.h>

void swap(void *vp1, void *vp2, size_t size) {
  void *buff = malloc(size);
  if (buff == NULL) {
    return;
  }
  memcpy(buff, vp1, size);
  memcpy(vp1, vp2, size);
  memcpy(vp2, buff, size);
  free(buff);
}
```
