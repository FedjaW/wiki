# Unfollow

If you can follow feeds, it only makes sense to be able to unfollow them.

## Assignment

1. Add a new SQL query to delete a feed follow record by user and feed url combination.
2. Add a new `unfollow` command that accepts a feed’s URL as an argument and unfollows it for the current user. This is, of course, a “logged in” command - use the new middleware.
