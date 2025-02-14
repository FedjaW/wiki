# Altering Tables

We often need to alter our database schema without deleting it and re-creating it. Imagine if Twitter deleted its database each time it needed to add a feature, that would be a disaster! Your account and all your tweets would be wiped out on a daily basis.

Instead, we can use the `ALTER TABLE` statement to make changes in place without deleting any data.

## ALTER TABLE

With SQLite an ALTER TABLE statement allows you to:

### 1. Rename a Table or Column

```sql
ALTER TABLE employees
RENAME TO contractors;

ALTER TABLE contractors
RENAME COLUMN salary TO invoice;
```

### 2. Add or DROP a Column

```sql
ALTER TABLE contractors
ADD COLUMN job_title TEXT;

ALTER TABLE contractors
DROP COLUMN is_manager;
```

## Assignment

We need to make some changes to the people table! At the moment, we have these six columns (shown as rows so we can display datatypes):

![img](03.png)

1. Rename the table to `users`
2. Rename the `tag` column to `username`.
3. Add the `password` (TEXT) column.

### SOLUTION

```sql
ALTER TABLE people
RENAME TO users;

ALTER TABLE users
RENAME COLUMN tag TO username;

ALTER TABLE users
ADD COLUMN password TEXT;
```