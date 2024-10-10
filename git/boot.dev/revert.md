# Revert

Where git 'reset' is a sledgehammer, git 'revert' is a scalpel.

A revert is effectively an anti commit. It does not remove the commit (like reset), but instead creates a new commit that does the exact opposite of the commit being reverted. It undoes the change but keeps a full history of the change and its undoing.

To revert a commit, you need to know the commit hash of the commit you want to revert. You can find this hash using git log.

```shell
git log
```

Once you have the hash, you can revert the commit using git revert.

```shell
git revert <commit-hash>
```
