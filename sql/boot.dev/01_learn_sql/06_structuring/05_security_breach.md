# Security Breach

Bad news! CashPal suffered its first security breach, one of the interns leaked a sensitive file that contained some user passwords. Even worse, the user table has passwords stored as plain text. Luckily, we know which passwords were leaked so we can warn these specific users before we move to a modern password storage system.

## Challenge

Write a query that returns the `name` and `username` for every user with a `password` equal to `backendDev`, `welovebootdev`, or `SQLrocks`. Order the records so that the names are in alphabetical order.

### Solution

```sql
SELECT name, username
FROM users
WHERE password IN ('backendDev', 'welovebootdev', 'SQLrocks')
ORDER BY name;
```