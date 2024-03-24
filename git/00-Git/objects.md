# Blob

under the hood git uses `git hash-object` to create a hash:

`echo 'Hello World' | git hash-object --stdin`
=> 557db03de997c86a4a028e1ebd3a1ceb225be238

generating the SHA1 of the contents, with metadata:

`echo 'blob 14\0Hello, World!' | openssl sha1`
=> 557db03de997c86a4a028e1ebd3a1ceb225be238

to print the type:
`git cat-file -t 557db03`
=> blob

to print the data:
`git cat-file -p 557db03` 
=> Hello World

# Tree

A tree contains pointers:

- to blobs
- to other trees

and metadata:

- type of pointer  (blob or tree)
- filename or directorie name
- mode (executable file, symbolic link, ...)

It represents the directories and subdirectories.

# Commit

a commit is code snapshot.

a commit points to:

- a tree

and contains metadata:

- author and commiter
- date
- message
- parent commit (one or more)

the SHA1 of the commit is the hash of all this information.

# References

- Tags
- Branches
- HEAD - a pointer to the current commit

`cat .git/HEAD` 
=> ref: refs/heads/main

`cat .git/refs/heads/main`
=> 480e1f4494552298703e750594792444516ec48b
Points to the initial commit ^

# Branch

- A branch ist just a pointer to a particular commit
- The pointer of the current branch changes as new commits are made

# HEAD

- **HEAD** is how git knows what branch you are currently on, and what the next parent will be
- It is a pointer
    - Is usually points at the **_name_** of the current branch
    - But it can point to any commit too (detached HEAD)
- It moves when:
    - You make a commit in the currently active branch
    - When you checkout a new branch

## Detached HEAD

- Sometimes you need to checkout a specific commit instead of a branch
- git moves HEAD pointer to that commit
- as soon as you checkout a different branch or commit, the value of HEAD will point to the new SHA
- There is no reference pointing to the commits you made in a detached head

To safe your work in a detached head

- Create a new branch that points to the last commit you made in a detached state
`git branch <newbranchname> <commit>`
- Why the last commit
    - Because the other commits point to their parents

Discard your work

- If you dont't point a new branch at those commits, they will no longer be referenced by git. (dangling commits)
- Eventually, the will be garbage collected. (every few weeks)
- There is a way to get them back, if they aren't garbage collected with reflog
