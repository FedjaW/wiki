# Joins

Joins are one of the most important features that SQL offers. Joins allow us to make use of the relationships we have set up between our tables. In short, joins allow us to query multiple tables at the same time.

## Inner Join

The simplest and most common type of join in SQL is the `INNER JOIN`. By default, a `JOIN` command is an `INNER JOIN`. An `INNER JOIN` returns all of the records in `table_a` that have matching records in `table_b` as demonstrated by the following Venn diagram.

![img](./01.png)

## On

To perform a table join, we need to tell the database how to "match up" the rows from each table. The `ON` clause specifies the columns from each table that should be compared.

When the same column name exists in both tables, we have to specify which table each column comes from using the table name (or an alias) followed by a dot `.` before the column name.

```sql
SELECT *
FROM employees
INNER JOIN departments
ON employees.department_id = departments.id;
```

In this query:

- `employees.department_id` refers to the `department_id` column from the employees table.
- `departments.id` refers to the `id` column from the `departments` table.

The `ON` clause ensures that rows are matched based on these columns, creating a relationship between the two tables.

The query above returns all the fields from both tables. The `INNER` keyword only affects the number of rows returned, not the number of columns. The `INNER JOIN` filters rows based on matching `department_id` and `id`, while the `SELECT *` ensures all columns from both tables are included.

## Why Is This Important?

In many databases, different tables might share the same column names, such as id. If you don't specify the table name (or alias) for a column, the database won't know which column to use for the join. For example, writing `ON id = id` won't work because the database can't distinguish between the id columns in each table.

## Assignment

Our frontend team is working on a profile page and would like to display a user's country name instead of just the country's two-letter code. Let's start by writing a simple join between the users table and countries table. We will expand on this query more in the next exercise.

- Write an `INNER JOIN` between `users` and `countries`
- Return all fields from both tables
- Join on the `country_code` field

### Solution

```sql
SELECT * FROM users
INNER JOIN countries
ON users.country_code = countries.country_code;
```