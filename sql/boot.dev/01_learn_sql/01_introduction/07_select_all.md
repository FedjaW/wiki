# Select All

Your boss was impressed that you got the query to work but now he's realized that he needs more than just the user names to send his thank you notes.

## Challenge

Alter the query so that it retrieves every field in the user records

### Solution

```sql
SELECT * FROM users;
```

#### up.sql

```sql
CREATE TABLE users (id INTEGER, name TEXT, age INTEGER, balance INTEGER, is_admin BOOLEAN);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (1, 'Branston Bryan', 28, 450.00, true);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (2, 'Skyler Black', 27, 200.00, true);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (3, 'Taylor Johnson', 33, 496.24, false);
INSERT INTO users (id, name, age, balance, is_admin) VALUES (4, 'Sydney Thompson', 33, 496.24, false);
```