# Reset

When doing end-to-end testing on a system like this, it’s always a bit annoying to have to manually reset the state of the database. For example, the `register` command fails if you run it again with the same username because the user already exists.

That’s the behavior we want, but we have to manually reset the database each time we run tests.

## Assignment

Let’s add a command to our CLI that will reset the database to a blank state. This will make development much easier for us (although probably not useful in production… maybe even disastrous).

1. Add a new query to the database that deletes all the users in the `users` table. Generate the Go code with `sqlc generate`. Use the `:exec` tag in the comment.
2. Add a new command called `reset` that calls the query. Report back to the user about whether or not it was successful with an appropriate exit code.