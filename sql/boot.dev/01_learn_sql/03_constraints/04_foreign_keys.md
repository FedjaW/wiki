# Foreign Keys

Foreign keys are what makes relational databases relational! Foreign keys define the relationships between tables. Simply put, a FOREIGN KEY is a field in one table that references another table's `PRIMARY KEY`.

## Creating a Foreign Key in SQLite

Creating a `FOREIGN KEY` in SQLite happens at table creation! After we define the table fields and constraints we add a named CONSTRAINT where we define the `FOREIGN KEY` column and its REFERENCES.

Here's an example:

```sql
CREATE TABLE departments (
    id INTEGER PRIMARY KEY,
    department_name TEXT NOT NULL
);

CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    department_id INTEGER,
    CONSTRAINT fk_departments
    FOREIGN KEY (department_id)
    REFERENCES departments(id)
);
```

In this example, an `employee` has a `department_id`. The `department_id` must be the same as the `id` field of a record from the `departments` table. `fk_departments` is the specified name of the constraint.

`CONSTRAINT fk_departments`: create a constraint called `fk_departments`

`FOREIGN KEY (department_id)`: make this constraint a foreign key assigned to the `department_id` field

`REFERENCES departments(id)`: link the foreign field `id` from the departments table

## Assignment

Our `users` table stores the country our users are from in a `country_code` field. We need some additional data about countries like their name, but we don't want to bloat our `users` table with all that country data.

The "locations" team at CashPal has created a "countries" table, and we can link a user to their country by setting a foreign key in the users table. Creating a foreign key without the CONSTRAINT keyword means the name of the constraint is auto-assigned.

Take a look at the code. There's an issue with the `INSERT` statements again! Fix up the data so no foreign key constraints are violated. You'll need to reference the setup in the `up.sql` tab in the code editor.

### Solution

Solutino was to fix a typo of a foreign_key constraint.