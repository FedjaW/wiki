# Count

We can use a `SELECT` statement to get a count of the records within a table. This can be very useful when we need to know how many records there are, but we don't particularly care what's in them.

Here's an example in SQLite:

```sql
SELECT COUNT(*) FROM employees;
```

The * in this case refers to a column name. We don't care about the count of a specific column - we want to know the number of total records so we can use the wildcard (*).

## Assignment

Here is the current state of our users table:

![img](./06.png)


Our business strategy team at CashPal wants to know how many users of the app we have. We can't use the `id` number to calculate the count because user accounts can be deleted!

Use a `COUNT(*)` statement to retrieve the number of records in the `users` table.

Note: In this course, stick to using `*` with `COUNT` unless the instructions specifically say to count a particular column.

### Solution

```sql
SELECT COUNT(*) FROM users;
```