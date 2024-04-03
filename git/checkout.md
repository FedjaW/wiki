# git checkout

```SHELL
# checkout a file from a specific commit
# will copy overwrite both working and staging area!
# this is a descruptive command!!! (no warning from git)
git checkout <commit> -- <file_path>
```

```SHELL
# restore a deleted file
# will copy overwrite both working and staging area!
# this is a descruptive command!!! (no warning from git)
git checkout <deleted_commit>^ -- <file_path>
```

The `--` signifies the end of a command operation and a start of pisitional parameters
