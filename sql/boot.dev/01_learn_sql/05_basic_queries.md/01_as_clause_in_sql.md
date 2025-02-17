# As Clause in SQL

Sometimes we need to structure the data we return from our queries in a specific way. An `AS` clause allows us to "alias" a piece of data in our query. The alias only exists for the duration of the query.

## As Keyword

The following queries return the same data:

```sql
SELECT employee_id AS id, employee_name AS name
FROM employees;
```

```sql
SELECT employee_id, employee_name
FROM employees;
```

The difference is that the results from the aliased query would have column names `id` and `name` instead of `employee_id` and `employee_name`.

## Assignment

A user has asked us to find all the transactions on their account from their grandma. We thought it would be fun to rename the `note` field to `birthday_message` because we noticed all the transactions from grandma are birthday messages.

Return the `amount` and the `note` field (renamed to `birthday_message`) from the transactions table where the `sender_id` is `10` (grandma).

### Solution

```sql
SELECT amount, note AS birthday_message FROM transactions
WHERE sender_id = 10;
```