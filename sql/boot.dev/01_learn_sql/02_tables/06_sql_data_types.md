# SQL Data Types

SQL as a language can support many different data types. However, the datatypes that your database management system (DBMS) supports will vary depending on the specific database you're using.

SQLite only supports the most basic types, and we're using SQLite in this course!

## SQLite Data Types

Let's go over the data types supported by SQLite: and how they are stored.

1. `NULL` - Null value.
2. `INTEGER` - A signed integer stored in 0,1,2,3,4,6, or 8 bytes.
3. `REAL` - Floating point value stored as an 64-bit IEEE floating point number.
4. `TEXT` - Text string stored using database encoding such as UTF-8
5. `BLOB` - Short for Binary large object and typically used for images, audio or other multimedia.
6. `BOOLEAN` - Boolean values are written in SQLite queries as `true` or `false`, but are recorded as `1` or `0`.

## Boolean Values

It's important to note - SQLite does not have a separate `BOOLEAN` storage class. Instead, boolean values are stored as integers:

- `0` = `false`
- `1` = `true`

It's not actually all that weird, `boolean` values are just binary bits after all!

SQLite will still let you write your queries using `boolean` expressions and `true/false` keywords, but it will convert the booleans to integers under-the-hood.