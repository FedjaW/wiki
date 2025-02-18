# Many to Many

A many-to-many relationship occurs when multiple records in one table can be related to multiple records in another table.

## Examples of Many-To-Many Relationships

A products table and a suppliers table - Products may have 0 to many suppliers, and suppliers can supply 0 to many products.
A classes table and a students table - Students can take potentially many classes and classes can have many students enrolled.

## Joining Table

Joining tables help define many-to-many relationships between data in a database. As an example when defining the relationship above between products and suppliers, we would define a joining table called products_suppliers that contains the primary keys from the tables to be joined.

Then, when we want to see if a supplier supplies a specific product, we can look in the joining table to see if the ids share a row.

## Unique Constraint Across 2 Fields

When enforcing specific schema constraints we may need to enforce the UNIQUE constraint across two different fields.

```sql
CREATE TABLE product_suppliers (
  product_id INTEGER,
  supplier_id INTEGER,
  UNIQUE(product_id, supplier_id)
);
```

This lets multiple rows share the same product_id or supplier_id, but it prevents any two rows from having both the same product_id and supplier_id.

## Assignment

We have another issue with our current user<->country relationship! In the current schema, a user can have many countries associated with it, but there are duplicate records! If two users are associated with the United States, we're creating two countries records.

It would be better if each country only had a single record. That way, when a country changes its metadata (for example, it decides to rename itself) we only have to update one record.

Use a joining table to de-deduplicate country data.

- Remove the user_id field from the countries table
- Create a new table called users_countries. It should have two fields:
  - country_id
  - user_id
- Add a "unique together" constraint on those id fields

### Solution

```sql
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_admin BOOLEAN
);

CREATE TABLE users_countries (
  country_id INTEGER,
  user_id INTEGER,
  UNIQUE(country_id, user_id)
);

CREATE TABLE countries (
  id INTEGER PRIMARY KEY,
  country_code TEXT,
  name TEXT
);
```