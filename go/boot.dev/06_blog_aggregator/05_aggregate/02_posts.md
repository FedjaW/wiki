# Posts

Time to actually store the posts in the database! We'll also add a `browse` command to view all the posts from the feeds the user follows, right in the terminal!

## Assignment

1. Add a `posts` table to the database.

A post is a single entry from a feed. It should have:

- `id` - a unique identifier for the post
- `created_at` - the time the record was created
- `updated_at` - the time the record was last updated
- `title` - the title of the post
- `url` - the URL of the post (this should be unique)
- `description` - the description of the post
- `published_at` - the time the post was published
- `feed_id` - the ID of the feed that the post came from

Some of these fields can probably be null, others you might want to be more strict about - it's up to you.

2. Add a "create post" SQL query to the database. This should insert a new post into the database.

3. Add a "get posts for user" SQL query to the database. Order the results so that the most recent posts are first. Make the number of posts returned configurable.

4. Update your scraper to save posts. Instead of printing out the titles of the posts, save them to the database!

- If you encounter an error where the post with that URL already exists, just ignore it. That will happen a lot.
- If it's a different error, you should probably log it.
- Make sure that you're parsing the "published at" time properly from the feeds. Sometimes they might be in a different format than you expect, so you might need to handle that.
- You may have to manually convert the data into database/sql types.

5. Add the `browse` command. It should take an optional "limit" parameter. If it's not provided, default the limit to 2.
