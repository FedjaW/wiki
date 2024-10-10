# Revert vs Reset

git reset --soft: Undo commits but keep changes staged
git reset --hard: Undo commits and discard changes
git revert: Create a new commit that undoes a previous commit

## When to reset

If you're working on your own branch, and you're just undoing something you've already committed, say you're cleaning everything up so you can open a pull request, then git reset is probably what you want.
