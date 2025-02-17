# Underscore Operator

As discussed, the `%` wildcard operator matches zero or more characters. Meanwhile, the `_` wildcard operator only matches a single character.

```sql
SELECT * FROM products
    WHERE product_name LIKE '_oot';
```

The query above matches products like:

- boot
- root
- foot

```sql
SELECT * FROM products
    WHERE product_name LIKE '__oot';
```

The query above matches products like:

- shoot
- groot

## Assignment

HR has been able to narrow down their query further! They want a report of all user data for users whose names start with `Al` and are exactly 5 characters long.

### Solution 

```sql
SELECT * FROM users
WHERE name LIKE 'Al___';
```