# Diff

The git diff command shows you the differences between... stuff. Differences between commits, the working tree, etc.

I frequently use it to look at the changes between the current state of my code and the last commit. For example:


```shell
# show the changes between the working tree and the last commit
git diff
```

```shell
# show the differences between the previous commit and the current state, including the last commit and uncommitted changes
git diff HEAD~1
```

```shell
# show the change between two commits
git diff COMMIT_HASH_1 COMMIT_HASH_2
```
