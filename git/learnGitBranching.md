# Learn Git Branching

Source: <https://learngitbranching.js.org/>

## Start

- Moving upwards one commit at a time with `^`
- Moving upwards a number of times with `~<num>`

## The "^" operator

Let's look at the Caret `^` operator first. Each time you append that to a ref name, you are telling Git to find the parent of the specified commit.

So saying `main^` is equivalent to _the first parent of main_.

`main^^` is the grandparent (second-generation ancestor) of main.

## The "~" operator

Say you want to move a lot of levels up in the commit tree. It might be tedious to type `^` several times, so Git also has the tilde `~` operator.

The tilde operator (optionally) takes in a trailing number that specifies the number of parents you would like to ascend.

## Branch forcing

One of the most common ways to use relative refs is to move branches around. You can directly reassign a branch to a commit with the -f option. So something like:

`git branch -f main HEAD~3`

moves (by force) the main branch to three parents behind HEAD.

**Note:** In a real git environment `git branch -f <command>` is not allowed for your current branch.

## Reversing Changes in Git

There are two primary ways to undo changes in Git -- one is using `git reset` (see: [[reset.md]]) and the other is using `git revert`.

### Git Reset

`git reset` reverses changes by moving a branch reference backwards in time to an older commit. In this sense you can think of it as "rewriting history;" `git reset` will move a branch backwards as if the commit had never been made in the first place.

### Git Revert

While resetting works great for local branches on your own machine, its method of "rewriting history" doesn't work for remote branches that others are using.

In order to reverse changes and share those reversed changes with others, we need to use `git revert`.

## Git Cherry-Pick

The first command in this series is called git cherry-pick. It takes on the following form:

- `git cherry-pick <Commit1> <Commit2> <...>`

It's a very straightforward way of saying that you would like to copy a series of commits below your current location (HEAD).

## Git Interactive Rebase

Git cherry-pick is great when you know which commits you want (and you know their corresponding hashes) -- it's hard to beat the simplicity it provides.

But what about the situation where you don't know what commits you want? Thankfully git has you covered there as well! We can use interactive rebasing for this -- it's the best way to review a series of commits you're about to rebase.

All interactive rebase means Git is using the rebase command with the `-i` option.

If you include this option, git will open up a UI to show you which commits are about to be copied below the target of the rebase. It also shows their commit hashes and messages, which is great for getting a bearing on what's what.

## Git Amend

Here's another situation that happens quite commonly. You have some changes (newImage) and another set of changes (caption) that are related, so they are stacked on top of each other in your repository (aka one after another).

The tricky thing is that sometimes you need to make a small modification to an earlier commit. In this case, design wants us to change the dimensions of newImage slightly, even though that commit is way back in our history!!

We will overcome this difficulty by doing the following:

- We will re-order the commits so the one we want to change is on top with `git rebase -i`
- We will `git commit --amend` to make the slight modification
Then we will re-order the commits back to how they were previously with `git rebase -i`
- Finally, we will move main to this updated part of the tree to finish the level (via the method of your choosing)

There are many ways to accomplish this overall goal, but for now let's focus on this technique.

## Git Remote Branches

After `git clone` A new branch appeared in our local repository called `o/main`. This type of branch is called a `remote branch`; remote branches have special properties because they serve a unique purpose.

Remote branches reflect the state of remote repositories (since you last talked to those remote repositories). They help you understand the difference between your local work and what work is public -- a critical step to take before sharing your work with others.

Remote branches have the special property that when you check them out, you are put into detached HEAD mode. Git does this on purpose because you can't work on these branches directly; you have to work elsewhere and then share your work with the remote (after which your remote branches will be updated).

**To be clear:** Remote branches are on your local repository, not on the remote repository.

## What is o/?

You may be wondering what the leading `o/` is for on these remote branches. Well, remote branches also have a (required) naming convention -- they are displayed in the format of:

- `<remote name>/<branch name>`

Hence, if you look at a branch named `o/main`, the branch name is `main` and the name of the remote is `o`.

Most developers actually name their **main remote origin, not o**. This is so common that git actually sets up your remote to be named origin when you git clone a repository.

Unfortunately the full name of origin does not fit in our UI, so we use o as shorthand :( Just remember when you're using real git, your remote is probably going to be named `origin`!

## What fetch does

git fetch performs two main steps, and two main steps only. It:

- downloads the commits that the remote has but are missing from our local repository, and...
- updates where our remote branches point (for instance, o/main)
git fetch essentially brings our local representation of the remote repository into synchronization with what the actual remote repository looks like (right now).

## What fetch doesn't do

`git fetch`, however, does not change anything about your local state. It will not update your main branch or change anything about how your file system looks right now.

This is important to understand because a lot of developers think that running git fetch will make their local work reflect the state of the remote. It may download all the necessary data to do that, but it does not actually change any of your local files.
