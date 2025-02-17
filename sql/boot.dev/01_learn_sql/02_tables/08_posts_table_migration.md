# Posts Table Migration

CashPal chatter is a huge hit! After several weeks of use, the engineers at CashPal have decided that some changes need to be made to our posts table. You've been asked to write an up migration to alter the table.

## Challenge

Write an up migration for the posts table that achieves the following:

- The `author_id` column should be renamed to `poster_id`
- Add a new column named `is_edited` with a `BOOLEAN` type
- `DROP` the `is_sponsored` column

### Solution

```sql
ALTER TABLE posts 
RENAME COLUMN author_id TO poster_id;

ALTER TABLE posts 
ADD COLUMN is_edited BOOLEAN;

ALTER TABLE posts 
DROP COLUMN is_sponsored;


-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('posts');
```