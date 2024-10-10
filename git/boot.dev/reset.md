# Reset

## Git Reset Soft

The git reset command can be used to undo the last commit(s) or any changes in the index (staged but not committed changes) and the worktree (unstaged and not committed changes).

```shell
git reset --soft COMMITHASH
```

The --soft option is useful if you just want to go back to a previous commit, but keep all of your changes. Committed changes will be uncommitted and staged, while uncommitted changes will remain staged or unstaged as before.

## Git Reset Hard

In the last lesson, we undid a commit but kept the changes. We don't want to keep the changes to titles.md, here's how to reset those changes.

```shell
git reset --hard COMMITHASH
```
