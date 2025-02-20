# Performance Tuning

Some of CashPal's users are complaining that certain pages on the site have become painfully slow.

Your manager has done some investigation and found that a `JOIN` to the countries table is causing some slowness. They've determined that denormalizing country data is a necessary risk for some performance benefits.

## Assignment

Run a migration to speed up future performance:

1. Add `country_code` and `country_name` columns to the `users` table (these are not optional fields, you'll need a default).
2. Update the `users` table to copy the data from `countries` table to the `users` table's `country_code` and `country_name` columns.
3. Drop the `country_id` column from the `users` table.

### Bonus:

4. Index the `country_code` column (`country_code_idx`) so that users can quickly be looked up by country.
5. Get rid of the countries table.

### Tips

Here's some reminders on how to update a row and how to write a subquery.

```sql
UPDATE orders
SET product_id = 25, delivery_address = '124 Conch St., Bikini Bottom, Pacific Ocean'
WHERE id = 1;
```

```sql
SELECT name
FROM products
WHERE id = (SELECT product_id FROM orders WHERE id = 1);
```

### Solution

```sql
ALTER TABLE users ADD COLUMN country_name TEXT NOT NULL DEFAULT 'Wonderland';
ALTER TABLE users ADD COLUMN country_code TEXT NOT NULL DEFAULT 'WL';
CREATE INDEX country_code_idx ON users (country_code);
UPDATE users
SET country_name = (SELECT name FROM countries WHERE id = country_id),
    country_code = (SELECT country_code FROM countries WHERE id = country_id);
ALTER TABLE users DROP COLUMN country_id;
DROP TABLE countries;
```