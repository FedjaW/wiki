# Config

Gator is a multi-user CLI application. There’s no server (other than the database), so it’s only intended for local use, but just like games in the 90’s and early 2000’s, that doesn’t mean we can’t have multiplayer functionality on a single device!

We’ll use a single JSON file to keep track of two things:

1. Who is currently logged in
2. The connection credentials for the PostgreSQL database

The JSON file should have this structure (when prettified):

```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```

Note: There’s no user-based authentication for this app. If someone has the database credentials, they can act as any user. We’ll cover auth in other courses. We want to focus on SQL, CLIs and long-running services for this project.

## Assignment

Manually create a config file in your home directory, `~/.gatorconfig.json`, with the following content:

```json
{
  "db_url": "postgres://example"
}
```

Don’t worry about adding `current_user_name`, that will be set by the application.

2. Initialize a new Go module, `go mod init`. Create a new file, `main.go`. Add a `main` function. It doesn’t need to do anything yet.
3. Create the `internal` and `internal/config` directories to contain the `config` internal package. This package will be responsible for reading and writing the JSON file.

This package should have the following functionality exported so the main package can use it:

- Export a `Config` struct that represents the JSON file structure, including struct tags.
- Export a `Read` function that reads the JSON file found at `~/.gatorconfig.json` and returns a `Config` struct. It should read the file from the HOME directory, then decode the JSON string into a new `Config` struct. I used `os.UserHomeDir` to get the location of HOME.
- Export a `SetUser` method on the `Config` struct that writes the config struct to the JSON file after setting the `current_user_name` field.

I also wrote a few non-exported helper functions and added a constant to hold the filename.

- `getConfigFilePath() (string, error)`
- `write(cfg Config) error`
- `const configFileName = ".gatorconfig.json"`

But you can implement the internals of the package however you like.

4. Update the `main` function to:

  - Read the config file.
  - Set the current user to “lane” (actually, you should use your name instead) and update the config file on disk.
  - Read the config file again and print the contents of the config struct to the terminal.