# List Users

We need a users command that lists all the users and indicates which one is currently logged in.

## Assignment

1. dd a `GetUsers` query that returns all the users from the database. Generate the Go code with `sqlc generate`. Use the `:many` tag in the comment.
2. Add a new command called `users` that calls `GetUsers` and prints all the users to the console.

Print them in this format:

```txt
* lane
* allan (current)
* hunter
```

Ensure that the `(current)` follows the currently logged-in user.