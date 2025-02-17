# Discount Program

CashPal has decided to start two new discount programs targeting certain customers. The discounts are as follows:

1. All users over the age of 55 will qualify for a senior discount
2. In celebration of Canada Day - all Canadian (`country_code` "CA") users qualify for a discount

We want to return all the records from the `users` table, but we want to include an extra column to know whether each user qualifies for any of these discounts.

## Challenge

Return all the data from the users table and a temporary column aliased to `discount_eligible`.

The `discount_eligible` column should have a boolean value of `true` or `false` depending on whether the user matches any discount conditions listed above.

### Solution

```sql
SELECT *, IIF(age > 55 OR country_code = 'CA', true, false) AS discount_eligible
FROM users;
```