# Goose Migrations

Goose is a database migration tool written in Go. It runs migrations from a set of SQL files, making it a perfect fit for this project (we wanna stay close to the raw SQL).

## What Is a Migration?

A migration is just a set of changes to your database table. You can have as many migrations as needed as your requirements change over time. For example, one migration might create a new table, one might delete a column, and one might add 2 new columns.

An “up” migration moves the state of the database from its current schema to the schema that you want. So, to get a “blank” database to the state it needs to be ready to run your application, you run all the “up” migrations.

If something breaks, you can run one of the “down” migrations to revert the database to a previous state. “Down” migrations are also used if you need to reset a local testing database to a known state.

## Assignment

1. Install Goose

Goose is just a command line tool that happens to be written in Go. I recommend installing it using `go install`:

```cli
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run `goose -version` to make sure it’s installed correctly.

2. Create a `users` migration in a new `sql/schema` directory.

A “migration” in Goose is just a .sql file with some SQL queries and some special comments. Our first migration should just create a users table. The simplest format for these files is:

```txt
number_name.sql
```

For example, I created a file in sql/schema called 001_users.sql with the following contents:


```cli
-- +goose Up
CREATE TABLE ...

-- +goose Down
DROP TABLE users;
```

Write out the `CREATE TABLE` statement in full, I left it blank for you to fill in. A `user` should have 4 fields:

- `id`: a `UUID` that will serve as the primary key
- `created_at`: a `TIMESTAMP` that can not be null
- `updated_at`: a `TIMESTAMP` that can not be null
- `name`: a unique string that can not be null

The `-- +goose Up` and `-- +goose Down` comments are case sensitive and required. They tell Goose how to run the migration in each direction.


3. Get your connection string. A connection string is just a URL with all of the information needed to connect to a database. The format is:

```txt
protocol://username:password@host:port/database
```

Here are examples:

- macOS (no password, your username): `postgres://wagslane:@localhost:5432/gator`
- Linux (password from last lesson, postgres user): `postgres://postgres:postgres@localhost:5432/gator`

Test your connection string by running `psql`, for example:

```cli
psql "postgres://wagslane:@localhost:5432/gator"
```

It should connect you to the gator database directly. If it’s working, great. exit out of psql and save the connection string.

4. Run the up migration.

`cd` into the `sql/schema` directory and run:

```cli
goose postgres <connection_string> up
```

Run your migration! Make sure it works by using psql to find your newly created users table:


```cli
psql gator
\dt
```

5. Run the down migration to make sure it works (it should just drop the table). When you’re satisified, run the up migration again to recreate the table.

6. Add the connection string to the .gatorconfig.json file instead of the example string we used earlier. When using it with goose, you’ll use it in the format we just used. However, here in the config file it needs an additional sslmode=disable query string:

```txt
protocol://username:password@host:port/database?sslmode=disable
```

Your application code needs to know to not try to use SSL locally.