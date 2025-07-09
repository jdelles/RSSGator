# RSSGator

RSSGator is a command-line RSS feed aggregator written in Go.

## Prerequisites

- **PostgreSQL**: You must have a running PostgreSQL instance.
- **Go**: Install Go (version 1.18 or newer recommended).

## Installation

Install the `gator` CLI using:

```sh
go install
```

This will place the `gator` binary in your `$GOPATH/bin` or `$HOME/go/bin`.

## Configuration

Create a configuration file at `~/.gatorconfig.json` with the following structure:

```json
{
  "db_url": "postgres://youruser:yourpass@localhost:5432/gator?sslmode=disable",
  "current_user_name": "yourusername"
}
```

- Replace `youruser`, `yourpass`, and `yourusername` with your actual PostgreSQL credentials and desired username.

## Database Setup

Run the migrations to set up the database schema:

```sh
goose -dir sql/schema postgres "your-db-url" up
```

## Usage

Run the CLI from your terminal:

```sh
gator <command> [arguments]
```

### Example Commands

- **register `<username>`**  
  Register a new user.

- **login `<username>`**  
  Login as an existing user.

- **addfeed `<feed_name>` `<feed_url>`**  
  Add a new RSS feed and follow it.

- **follow `<feed_url>`**  
  Follow an existing feed by URL.

- **browse [limit]**  
  Browse recent posts (optionally specify how many to show).

- **agg `<duration>`**  
  Continuously aggregate feeds every `<duration>` (e.g., `1m`, `10s`).

## Future Development Ideas: 

- Add sorting and filtering options to the browse command

- Add pagination to the browse command

- Add concurrency to the agg command so that it can fetch more frequently

- Add a search command that allows for fuzzy searching of posts

- Add bookmarking or liking posts

- Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)

- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely

- Write a service manager that keeps the agg command running in the background and restarts it if it crashes

---

Enjoy using RSSGator!