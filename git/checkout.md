# git checkout

```SHELL
# checkout a file from a specific commit
# will copy overwrite both working and staging area!
# this is a disruptive command!!! (no warning from git)
git checkout <commit> -- <file_path>
```

```SHELL
# restore a deleted file
# will copy overwrite both working and staging area!
# this is a disruptive command!!! (no warning from git)
git checkout <deleted_commit> -- <file_path>
```

The `--` signifies the end of options and the start of positional parameters.

---

From Stackoverflow:
<https://unix.stackexchange.com/questions/11376/what-does-double-dash-double-hyphen-mean>

More precisely, a double dash (`--`) is used in most Bash built-in commands and many other commands to signify the end of command options, after which only positional ("non-option") arguments are accepted.
Example use: Let's say you want to grep a file for the string -v. Normally -v will be considered the option to reverse the matching meaning (only show lines that do not match), but with -- you can grep for the string -v like this:

```shell
grep -- -v file
```

## How to Checkout a File from Another Branch

```shell
git checkout <other-branch-name> -- path/to/your/folder
```
