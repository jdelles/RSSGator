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

---

Enjoy using RSSGator!