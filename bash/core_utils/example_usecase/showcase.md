Let's say we want to count how many occurrences of each word exists within a file.

```SHELL
cat subtitles_dutch.md | xargs | tr ' ' '\n' | sort | uniq -c | sort -nr
```
