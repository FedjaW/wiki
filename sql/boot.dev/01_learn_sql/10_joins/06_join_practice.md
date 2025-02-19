# Join Practice

Joins take some time to get used to, but the key to understanding them and using them effectively is practice!

## Multiple Joins

To incorporate data from more than two tables, you can utilize multiple joins to execute more complex queries!

```sql
SELECT *
FROM employees
LEFT JOIN departments
ON employees.department_id = departments.id
INNER JOIN regions
ON departments.region_id = regions.id;
```

## Assignment

Our front-end team is finalizing the profile page for CashPal. We need to write a query that returns all the `user` data they need for an individual user's profile. The query needs to return the following fields:

1. The user's `id`
2. The user's `name`
3. The user's `age`
4. The user's `username`
5. The user's country name, renamed to `country_name`
6. The sum of the successful transactions from the user, renamed to `balance`

Return only a single user record - specifically the one with `id=6`.

### Solution

```sql
SELECT users.id,
       users.name,
       users.age,
       users.username,
       countries.name as country_name,
       SUM(transactions.amount) as balance
FROM users
LEFT JOIN transactions
ON transactions.user_id = users.id
INNER JOIN countries
ON users.country_code = countries.country_code
WHERE users.id = 6 AND transactions.was_successful;
```