# ORDER BY

SQL also offers us the ability to sort the results of a query using `ORDER BY`. By default, the `ORDER BY` keyword sorts records by the given field in ascending order, or `ASC` for short. However, `ORDER BY` does support descending order as well with the keyword `DESC`.

## Examples

This query returns the `name`, `price`, and `quantity` fields from the products table sorted by `price` in ascending order:

```sql
SELECT name, price, quantity FROM products
    ORDER BY price;
```

This query returns the `name`, `price`, and `quantity` of the products ordered by the quantity in descending order:

```sql
SELECT name, price, quantity FROM products
    ORDER BY quantity DESC;
```

## Assignment

Write a query that lists all the records in the `transactions` table where:

- `amount` is `BETWEEN` 10 and 80 dollars.
- The results are sorted by `amount` in descending order.

### Solution

```sql
SELECT * FROM transactions
WHERE amount BETWEEN 10 AND 80
ORDER BY amount DESC;
```

Same as 

```sql
SELECT * FROM transactions
WHERE (amount >= 10 AND amount <= 80)
ORDER BY amount DESC;
```

not that `BETWEEN` is icluding 10 and 80