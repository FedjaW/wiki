# Logic

## Chaining commands with list operators

A list is when oyu put one or more commands on a given line.

List operators:

- ;
- &
- &&
- ||

& lets you run commands asynchronously.

```SHELL
echo 123 & echo 456
# [1] 74 -> echo 123 was send to the background
# 456
# 123
```

`;` lets you run commands sequentialy.

```SHELL
echo 123 ; echo 456
# 123
# 456
```

```SHELL
# will be executed even if error appears
ls non_existing_dir ; echo 123
# ls: cannot access...
# 123
```

That brings us to the next operators to chain commands:

The `&&` (and) operator makes it so that the second command only runs if the first one was successful.

The `||` (or) operator makes it so that the second command only runs if the first one failed.
