# What Is SQL?

Structured Query Language, or SQL (pronounced "squeel" by the in-crowd), is the primary programming language used to manage and interact with relational databases. SQL can perform various operations such as creating, updating, reading, and deleting records within a database.

## CashPal

In this course, we will be building the database for a pretend PayPal clone called CashPal! Storing information related to people's money, transactions, and identity is very important! So we will need to make sure we use proper conventions to build a safe, and reliable database architecture that our users can rely on.

## Assignment

I have provided a simple SQL statement for you that retrieves some records from a table. However there isn't a `people` table, the table in our database is called `users`. Fix the bug by changing `people` to `users` within the `SELECT` statement.

### Solution

```sql
SELECT * FROM users;
```

### up.sql

```sql
CREATE TABLE users (id INTEGER, name TEXT, age INTEGER, is_admin BOOLEAN);
INSERT INTO users (id, name, age, is_admin) VALUES (1, 'John Doe', 27, false);
INSERT INTO users (id, name, age, is_admin) VALUES (2, 'Sally Rae', 18, true);

```