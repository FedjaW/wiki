# What is an Algorithm?

In the context of computer science, an algorithm is a finite sequence of well-defined, computer-implementable instructions. In short, an algorithm is:

- Defined: there is a specific sequence of steps that performs a task
- Unambiguous: there is a "correct" and "incorrect" interpretation of the steps
- Implementable: it can be executed using software and hardware

Calculate the median:

```py
def median(nums):
    if len(nums) == 0:
        return None
    nums = sorted(nums)
    n = len(nums)
    if n % 2 == 0:
        return (nums[n // 2 - 1] + nums[n // 2]) / 2
    return nums[n // 2]
```
