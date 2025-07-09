package main

import (
	"database/sql"
	"fmt"
	"os"

	"RSSGator/commands"
	"RSSGator/handlers"
	"RSSGator/internal/config"
	"RSSGator/internal/database"
	"RSSGator/middleware"
	_ "github.com/lib/pq"
)

func main() {
	configPointer, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %v\n", err)
		os.Exit(1)
	}
	dbURL := configPointer.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	appState := commands.State{
		Cfg: configPointer,
		Db:  dbQueries,
	}
	cmds := &commands.Commands{
		Handlers: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerUsers)
	cmds.Register("agg", handlers.HandlerAgg)
	cmds.Register("addfeed", middleware.MiddlewareLoggedIn(handlers.HandlerAddFeed))
	cmds.Register("feeds", handlers.HandlerFeeds)
	cmds.Register("follow", middleware.MiddlewareLoggedIn(handlers.HandlerFollow))
	cmds.Register("following", middleware.MiddlewareLoggedIn(handlers.HandlerFollowing))
	cmds.Register("unfollow", middleware.MiddlewareLoggedIn(handlers.HandlerUnfollow))
	cmds.Register("browse", middleware.MiddlewareLoggedIn(handlers.HandlerBrowse))
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a valid command")
		os.Exit(1)
	}
	cmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}
	if err := cmds.Run(&appState, cmd); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
