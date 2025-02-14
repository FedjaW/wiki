# Broken Query

Your boss at CashPal wants to get a list of names for all the users in the database so he can thank them for joining. After hours of trying, he is insistent that the database must be broken and has asked you to take a look at his query.

## Challenge

Fix the error in the query so that it returns all the user names from the database.

### Solution

```sql
SELECT name FROM users;
```

#### up.sql

```sql
CREATE TABLE users (id INTEGER, name TEXT, age INTEGER, balance INTEGER, is_admin BOOLEAN);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (1, 'Branston Bryan', 28, 450.00, true);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (2, 'Skyler Black', 27, 200.00, true);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (3, 'Taylor Johnson', 33, 496.24, false);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (4, 'Sydney Thompson', 33, 496.24, false);
```
