# Banking Relationships

Someone added new columns to CashPal's users table that include information on the user's banking institution. Unfortunately, the developer that made these changes had assumed that a user could only have a single banking institution.

We know that a user can be a member of multiple banks and banks can service a great deal of CashPal users. You've been tasked with cleaning up the users table and creating any additional tables to properly support a many-to-many relationship between banks and users.

## Challenge

Given what you've learned about normalization and many-to-many relationships, do the following:

- Create a new table called `banks` that contains any bank-related columns that were incorrectly added to the `users` table and were prefixed with `bank_`. (Don't worry about altering the users table right now.)
- Create a joining table named `users_banks` with `user_id` and `bank_id` columns. Add any constraints such that there is never a duplicate row with the same `user_id` and bank_id combination.

### Solution

```sql
CREATE TABLE banks (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  routing_number INTEGER NOT NULL
);

CREATE TABLE users_banks (
  user_id INTEGER,
  bank_id INTEGER,
  UNIQUE(user_id, bank_id)
);
```