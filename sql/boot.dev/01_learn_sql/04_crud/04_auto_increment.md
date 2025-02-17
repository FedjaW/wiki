# Auto Increment

Many dialects of SQL support an AUTO INCREMENT feature. When inserting records into a table with AUTO INCREMENT enabled, the database will assign the next value automatically. In SQLite an integer id field that has the PRIMARY KEY constraint will auto increment by default!

## Id's

Depending on how your database is set up, you may be using traditional ids or you may be using UUIDs. SQL doesn't support auto incrementing a uuid so if your database is using them your server will have to handle the changing uuid's for each record.

## Using AUTO INCREMENT in SQLite

We are using traditional ids in our database, so we can take advantage of the auto increment feature. Different dialects of SQL will implement this feature differently, but in SQLite any column that has the INTEGER PRIMARY KEY constraint will auto increment! So we can omit the id field within the INSERT statement and allow the database to automatically add that field for us!

## Assignment

Let's add some more records into our users table, but allow the database to automatically increment the id field. Add the following records to the database:

### Record 1

- name: Lance
- age: 20
- country_code: US
- username: LanChr
- password: bootdevisbest
- is_admin: false

### Record 2

- name: Tiffany
- age: 28
- country_code: US
- username: Tifferoon
- password: autoincrement
- is_admin: true


### Solution

```sql
INSERT INTO users(name, age, country_code, username, password, is_admin)
VALUES ('Lance', 20, 'US', 'LanChr', 'bootdevisbest', false);
INSERT INTO users(name, age, country_code, username, password, is_admin)
VALUES ('Tiffany', 28, 'US', 'Tifferoon', 'autoincrement', true);

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT * FROM users;
```