Sum each line and display the max sum

```SHELL
cat numbers.md | tr ' ' '+' | bc | sort -nr | head -1
```