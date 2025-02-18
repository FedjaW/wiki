# No Tables

When working on a back-end application, this doesn't come up often, but it's important to remember that SQL is a full programming language. We usually use it to interact with data stored in tables, but it's quite flexible and powerful.

For example, you can SELECT information that's simply calculated, with no tables necessary.

```sql
SELECT 5 + 10 as sum;
```

## Assignment

Finance has found that people who have lived longer than 40 years need to start thinking about retirement. Write a query that returns all columns of all the `users` who are more than 40 years old. Unfortunately, this table awkwardly stores age in days in the `age_in_days` field.

Use a subquery to convert `40 years -> days` and filter on that. Assume every year has 365 days.

### Solution

```sql
SELECT * FROM users
WHERE age_in_days > (
  SELECT 365 * 40
)
```

SAME AS

```sql

SELECT * FROM users
  WHERE (
    SELECT age_in_days/365 as age
  ) > 40;
```