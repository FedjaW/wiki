# Postgres

PostgreSQL is a production-ready, open-source database. It’s a great choice for many web applications, and as a back-end engineer, it might be the single most important database to be familiar with.

## How Does PostgreSQL Work?

Postgres, like most other database technologies, is itself a server. It listens for requests on a port (Postgres’ default is `:5432`), and responds to those requests. To interact with Postgres, first you will install the server and start it. Then, you can connect to it using a client like psql or PGAdmin.

1. Install Postgres v15 or later.

### macOS with brew

```cli
brew install postgresql@15
```

### Linux / WSL (Debian). Here are the docs from Microsoft (<https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql>), but simply:

```cli
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. Ensure the installation worked. The `psql` command-line utility is the default client for Postgres. Use it to make sure you’re on version 15+ of Postgres:

```cli
psql --version
```

3. (Linux only) Update postgres password:

```cli
sudo passwd postgres
```

Enter a password, and be sure you won’t forget it. You can just use something easy like postgres.

4. Start the Postgres server in the background

- Mac: `brew services start postgresql`
- Linux: `sudo service postgresql start`

5. Connect to the server. I recommend simply using the `psql` client. It’s the “default” client for Postgres, and it’s a great way to interact with the database. While it’s not as user-friendly as a GUI like PGAdmin, it’s a great tool to be able to do at least basic operations with.

Enter the psql shell:

- Mac: `psql postgres`
- Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:

```cli
postgres=#
```

6. Create a new database. I called mine `gator`:

```cli
CREATE DATABASE gator;
```

7. Connect to the new database:

```cli
\c gator
```

You should see a new prompt that looks like this:

```cli
gator=#
```

8. Set the user password (Linux only)

```cli
ALTER USER postgres PASSWORD 'postgres';
```

For simplicity, I used `postgres` as the password. Before, we altered the system user’s password, now we’re altering the database user’s password.

9. Query the database

From here you can run SQL queries against the gator database. For example, to see the version of Postgres you’re running, you can run:

```cli
SELECT version();
```

You can type exit to leave the psql shell. Run and submit the tests using the CLI.