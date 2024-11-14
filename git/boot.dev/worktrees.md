# Worktrees

## What is a worktree?

A worktree (or "working tree" or "working directory") is just the directory on your filesystem where the code you're tracking with Git lives. Usually, it's just the root of your Git repo (where the .git directory is). It contains:

Tracked files (files that Git knows about)
Untracked files (files that Git doesn't know about)
Modified files (files that Git knows about that have been changed since the last commit)

## The worktree command

Git has the git worktree command that allows us to work with worktrees. The first subcommand we'll worry about is simple:

```shell
git worktree list
```

## Linked Worktrees

We've talked about:

1. Stash (temporary storage for changes)
2. Branches (parallel lines of development)
3. Clone (copying an entire repo)

Worktrees accomplish a similar goal (allow you to work on different changes without losing work), but are particularly useful when:

1. You want to switch back and forth between the two change sets without having to run a bunch of git commands (not branches or stash)
2. You want to keep a light footprint on your machine that's still connected to the main repo (not clone)

## Create a linked worktree

To create a new worktree at a given path:

```shell
git worktree add <path> [<branch>]
```

Note: <branch> is optional, it will use the last part of the path as the branch name.

## Delete Worktrees

You may never need to stash again! Okay, stash is still useful for tiny stuff, but worktrees are so much better for long-lived changes.

However, at some point you will need to clean up your worktrees. The simplest way is the remove subcommand:

```shell
git worktree remove WORKTREE_NAME
```

An alternative is to delete the directory manually, then prune all the worktrees (removing the references to deleted directories):

```shell
git worktree prune
```
