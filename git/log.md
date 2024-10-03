# git log

```SHELL
# will list all commits since yesterday
git log --since="yesterday"
```

```SHELL
# lists all commits of the file hello.txt even if renamed
git log --name-status --follow --oneline hello.txt
```

```SHELL
# demonstrates a nice combination
git log --grep="foo" --author=hans --since=2.weeks
```

```SHELL
# you can filter for (R) renamed
git log --diff-filter=R --find-renames
# or filter for (M) modified
git log --diff-filter=M --oneline
```

```SHELL
# A nice one to see the graph
git log --oneline --graph --all
```