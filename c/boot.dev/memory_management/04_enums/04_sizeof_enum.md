# Sizeof Enum

The same `sizeof` operator that we've talked about works on enums.

Generally, enums in C are the same size as an `int`. However, if an enum value exceeds the range of an `int`, the C compiler will use a larger integer type to accommodate the value, such as an unsigned `int` or a long.

- unsigned `int` doesn't represent negative numbers, so it can represent larger positive numbers.
- `long` is just a larger integer type than `int`, so it can represent larger numbers.

## Just Fancy Integers

Enums are often used to represent the possibilities in a set. For example:

- `SMALL` = 0
- `MEDIUM` = 1
- `LARGE` = 2
- `EXTRA_LARGE` = 3

Your code probably cares a lot about which size a variable represents, but it probably doesn't care that `SMALL` happens to be `0` under the hood. From the compiler's perspective, enums are just fancy integers.