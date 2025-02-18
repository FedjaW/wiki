# ORDER BY and Limit

When using both `ORDER BY` and `LIMIT`, the `ORDER BY` clause must come first.

## Assignment

An HR employee got into the Git repository where we store all the queries and tried to update one himself.

Fix the bug in the SQL query.

### Solution

```sql
SELECT * FROM transactions
WHERE amount BETWEEN 10 AND 80
ORDER BY amount DESC
LIMIT 4;
```

switched order by and limit statements so that order by comes first.