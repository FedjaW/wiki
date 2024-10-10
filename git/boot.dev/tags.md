# Tags

A tag is a name linked to a commit that doesn't move between commits, unlike a branch. Tags can be created and deleted, but not modified.

To list all current tags:

```shell
git tag
```

To create a tag on the current commit:

```shell
git tag -a "tag name" -m "tag message"
```

Note: A commit can have multiple tags
