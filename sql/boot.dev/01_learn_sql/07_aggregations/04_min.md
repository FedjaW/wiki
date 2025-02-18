# Min

The `min` function works the same as the `max` function but finds the lowest value instead of the highest value.

```sql
SELECT product_name, min(price)
FROM products;
```

This query returns the `product_name` and the `price` fields of the record with the lowest price.

## Assignment

Use a `min` aggregation to find only the `age` of our youngest CashPal user in the United States in the users table. The `country_code` of the United States is `US`.

## Users Table

| id | name | age | country_code | username | password | is_admin |

### Solution

```sql
SELECT min(age)
FROM users
WHERE country_code = 'US';
```