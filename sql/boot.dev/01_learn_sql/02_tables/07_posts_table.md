# Posts Table

You're working on CashPal's latest social media feature named "CashPal Chatter", which aims to revolutionize financial discussions online. Users can make a typical social media post that contains a picture of their latest purchase along with text and some other information. You've been assigned to create the new "posts" table.

## Challenge

Write an SQL statement to create a new table named `posts` which should contain the following columns:

- `id`
- `image_url`
- `description`
- `author_id`
- `is_sponsored`

Use data types that make the most sense given the column name. For ID columns, assume we can just use `INTEGER`.

### Solution

```sql
CREATE TABLE posts (
  id INTEGER,
  image_url TEXT,
  description TEXT,
  author_id INTEGER,
  is_sponsored BOOLEAN
);

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('posts');
```