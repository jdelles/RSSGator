package handlers

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
	"RSSGator/commands"
	"RSSGator/internal/database"
)

func HandlerRegister(s *commands.State, cmd commands.Command) error {
    if len(cmd.Args) == 0 {
        return fmt.Errorf("register expects a single argument, the username")
    }
    name := cmd.Args[0]
    userParams := database.CreateUserParams{
        ID:        uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Name:      name,
    }
    ctx := context.Background()
    user, err := s.Db.CreateUser(ctx, userParams)
    if err != nil {
        return fmt.Errorf("user with name '%s' already exists or could not be created: %w", name, err)
    }
    if err := s.Cfg.SetUser(name); err != nil {
        return fmt.Errorf("failed to set current user in config: %w", err)
    }
    fmt.Printf("User '%s' created successfully!\n", name)
    fmt.Printf("User data: %+v\n", user)
    return nil
}
