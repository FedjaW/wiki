# Average

Just like we may want to find the minimum or maximum values within a dataset, sometimes we need to know the average!

SQL offers us the `AVG()` function. Similar to `MAX()`, `AVG()` calculates the average of all non-NULL values.

```sql
SELECT avg(song_length)
FROM songs;
```

This query returns the average `song_length` in the `songs` table!

## Assignment

Our marketing team is trying to determine the best marketing channels to advertise through but they need more information about our current users. They wish to know the average age of users in the United States.

Return a single value representing the average age of all `users` where the `country_code` is `US`.

### Solution

```sql
SELECT avg(age)
FROM users
WHERE country_code = 'US';
```