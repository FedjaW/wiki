# One to Many

When talking about the relationships between tables, a one-to-many relationship is probably the most commonly used relationship.

A one-to-many relationship occurs when a single record in one table is related to potentially many records in another table.

Note: The one->many relation only goes one way, a record in the second table can not be related to multiple records in the first table!

## Examples of One-to-Many Relationships

- A customers table and an orders table. Each customer has 0, 1, or many orders that they've placed.
- A users table and a transactions table. Each user has 0, 1, or many transactions that they've taken part in.

```sql
CREATE TABLE customers (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    amount INTEGER NOT NULL,
    customer_id INTEGER,
    CONSTRAINT fk_customers
    FOREIGN KEY (customer_id)
    REFERENCES customers(id)
);
```

## Assignment

It turns out that when we originally designed the CashPal database schema we assumed that users would only have a single country they lived in. With digital nomads becoming a thing, it turns out many users have dual citizenship.

Instead of a single users table where each user has a single country_code, do the following:

- Remove the country_code field from the users table
- Create a new table called countries with 4 fields:
    - id: an integer primary key
    - country_code: a TEXT
    - name: a TEXT
    - user_id: an integer foreign key to the users table's id field

### Solution

```sql
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  -- country_code TEXT NOT NULL,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_admin BOOLEAN
);

CREATE TABLE countries (
  id INTEGER PRIMARY KEY,
  country_code TEXT,
  name TEXT,
  user_id INTEGER,
  CONSTRAINT fk_users
  FOREIGN KEY (user_id)
  REFERENCES users(id)
);
```