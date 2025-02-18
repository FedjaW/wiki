# Sum

The `SUM` aggregation function returns the sum of a set of values.

For example, the query below returns a single record containing a single field. The returned value is equal to the total salary being collected by all of the `employees` in the employees table.

```sql
SELECT SUM(salary)
FROM employees;
```

Which returns:

SUM(SALARY)
2483

## Assignment

We need to be able to calculate the current balance for a given user because we don't (yet) store the running balance on each individual transaction record.

Write a query that returns the `SUM` aggregation of the `amounts` for all of Bob's successful transactions (`user_id` is `9`).

### Solution

```sql
SELECT SUM(amount)
FROM transactions
WHERE user_id = 9 AND was_successful = true;
```