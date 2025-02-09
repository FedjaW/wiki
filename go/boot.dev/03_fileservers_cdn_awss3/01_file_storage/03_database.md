# Database

Tubely’s architecture is simple. We’re using:

- Golang to write the application
- SQLite as a database. SQLite is a traditional relational database that works out of a single flat file, meaning it doesn’t need a separate server process to run.
- Later, we’ll also use the filesystem and S3 to store large files

However, small structured data, like user records, will always live in the SQLite database.

## Assignment

1. With the server running, create a Tubely account by entering the following email and password and clicking “sign up”:

- `Email`: admin@tubely.com
- `Password`: password

We’ll use these credentials for the entire course!

2. Install SQLite 3. The application is already set up to use SQLite, I just want you to be able to use the CLI to manually inspect the database.

```cli
# linux
sudo apt update
sudo apt install sqlite3

# mac
brew update
brew install sqlite3
```

3. Run `sqlite3 tubely.db` to open the database file (that should have been created automatically by the server when you started it). Run a `select \* from users;` to see the users table. You should see yourself in there!
4. Type `.exit` to exit the SQLite CLI
