# Create a Foreign Key

We managed to fix the `banks` table but now we want to have the added safety of foreign key constraints on our `users` table. For this challenge, the users table contains a `bank_id`, we want to make sure that the user's `bank_id` references an `id` from the `banks` table.

## Challenge

Change the table creation statement to include a `FOREIGN KEY` constraint on the `users` table named `fk_banks` that `REFERENCES` an `ID` from the banks table.

### Solution

```sql
CREATE TABLE banks (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    location TEXT NOT NULL,
    routing_number INTEGER NOT NULL
);

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL,
    country_code TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOLEAN,
    bank_id INTEGER,
    CONSTRAINT fk_banks
    FOREIGN KEY (bank_id)
    REFERENCES banks(id)
);

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT sql AS users_table FROM sqlite_schema WHERE type='table' AND name='users';
```

Solution was to add

```sql
    CONSTRAINT fk_banks
    FOREIGN KEY (bank_id)
    REFERENCES banks(id)
```
