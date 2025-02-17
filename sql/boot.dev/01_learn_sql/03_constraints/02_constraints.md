# Constraints

A `constraint` is a rule we create on a database that enforces some specific behavior. For example, setting a `NOT NULL` constraint on a column ensures that the column will not accept `NULL` values.

If we try to insert a `NULL` value into a column with the `NOT NULL` constraint, the insert will fail with an error message. Constraints are extremely useful when we need to ensure that certain kinds of data exist within our database.

## Defining a NOT NULL Constraint

The `NOT NULL` constraint can be added directly to the `CREATE TABLE` statement.

```sql
CREATE TABLE employees(
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE,
    -- The UNIQUE constraint ensures that no two rows can have the same value in the 'name' column
    title TEXT NOT NULL
    -- The NOT NULL constraint ensures that the 'title' column cannot have NULL values
);
```

## SQLite Limitation

In other dialects of SQL you can `ADD CONSTRAINT` within an `ALTER TABLE` statement. SQLite does not support this feature so when we create our tables we need to make sure we specify all the constraints we want! Here's a list of SQL Features SQLite does not implement in case you're curious.

## Assignment

Thankfully all the tables we have created for CashPal up to this point have been for testing purposes! Now that we have a better understanding of constraints, let's rebuild our database with the proper constraints and tables!

Create the users table with the following fields and constraints:

- `id` - `INTEGER`, `PRIMARY KEY`
- `name` - `TEXT`, `NOT NULL`
- `age` - `INTEGER`, `NOT NULL`
- `country_code` - `TEXT`, `NOT NULL`
- `username` - `TEXT`, `UNIQUE`, `NOT NULL`
- `password` - `TEXT`, `NOT NULL`
- `is_admin` - `BOOLEAN`

### Solution

```sql
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  country_code TEXT NOT NULL,
  username  TEXT UNIQUE NOT NULL,
  password  TEXT NOT NULL,
  is_admin  BOOLEAN
);

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('users');

```