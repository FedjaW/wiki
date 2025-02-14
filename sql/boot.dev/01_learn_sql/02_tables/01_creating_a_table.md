# Creating a Table

To create a new table in a database, use the `CREATE TABLE` statement followed by the name of the table and the fields you want in the table.

```sql
CREATE TABLE employees (id INTEGER, name TEXT, age INTEGER, is_manager BOOLEAN, salary INTEGER);
```

Each field name is followed by its datatype. We'll get to data types in a minute.

It's also acceptable and common to break up the `CREATE TABLE` statement with some whitespace like this:

```sql
CREATE TABLE employees(
    id INTEGER,
    name TEXT,
    age INTEGER,
    is_manager BOOLEAN,
    salary INTEGER
);
```

## Assignment

Let's begin building a table for CashPal database! Create the `people` table with the following fields:

1. id - `Integer`
2. tag - `Text`
3. name - `Text`
4. age - `Integer`
5. balance - `Integer`
6. is_admin - `boolean`

### Tip

- The `PRAGMA TABLE_INFO(TABLENAME);` command returns information about a table and its fields. You don't need to edit this line, I just thought you might be curious!
- `INTEGER` and `INT` are slightly different data types. The tests expect `INTEGER`, not `INT`.

### Solution

```sql
CREATE TABLE people (
  id INTEGER,
  tag TEXT,
  name TEXT,
  age INTEGER,
  balance INTEGER,
  is_admin BOOLEAN
);
```