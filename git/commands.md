# Git - Commands

Add a file to the staging area

```shell
git add <filename>
```

Remove a file from the staging area

```shell
git rm <filename>
```

The a in `-am` will execute `git add .` automatically.

```shell
git commit -am "hello world!"
```

Will create an alias so you can call `git ac "hello world!"` and it will add and commit all files.

```shell
git config --global alias.ac "commit -am"
```

Will update the latest commit message without creating a new commit.

```shell
git commit --amend -m "new commit message"
```

Will include new changes in the latest commit and keeping the commit message. This will only work if you did not already push your commit to the remote repository.

```shell
git add .
git commit --amend -m --no-edit
```

Nicer log output of the commits:

```sh
git log --graph --oneline --decorate
```

Search for commit message

```sh
git log -S <searchterm>
```

Clean up untracked files or build artifacts:

```sh
git clean -df
```

Switch back to the previous branch:

```shell
git checkout -
```

List all branches:

```shell
git branch -a
```

Will create a new branch without switching to it

```shell
git branch <a_new_branch>
```

Crazy github feature:

```text
open up a github repo and press dot (.) on your keyboard, it will open up a vscode within the browser.
```

Stash local changes for later use:

```shell
git stash save coolstuff
git stash list
git stash apply 0
```

You can stash changes with a message.

```shell
git stash -m "my message"
```

Grab a single file from the stash

```shell
git checkout <stashname> -- <filename>
```

Remove all stashes

```shell
git stash clear
```

Inspect a stash (here 0 - the newest)

```shell
git stash show stash@{0}
```

Annotated Tag - Stores additional message

```shell
git tag -a v1.0 -m "Version 1.0 of my blob"
```

List all tags

```shell
git tag
```
