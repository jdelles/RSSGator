package handlers

import (
	"context"
	"fmt"
	"RSSGator/commands"
)

func HandlerLogin(s *commands.State, cmd commands.Command) error {
    if len(cmd.Args) == 0 {
        return fmt.Errorf("login expects a single argument, the username")
    }
    username := cmd.Args[0]
    ctx := context.Background()
    _, err := s.Db.GetUser(ctx, username)
    if err != nil {
        return fmt.Errorf("you can't login to an account that doesn't exist")
    }
    if err := s.Cfg.SetUser(username); err != nil {
        return fmt.Errorf("failed to set current user in config: %w", err)
    }
    fmt.Printf("The user has been set as %s\n", username)
    return nil
}