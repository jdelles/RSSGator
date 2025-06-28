package handlers

import (
	"context"
	"fmt"
	"time"
	
	"github.com/google/uuid"
	"RSSGator/commands"
	"RSSGator/internal/database"
)

func HandlerAddFeed(s *commands.State, cmd commands.Command) error {
    if len(cmd.Args) < 2 {
        return fmt.Errorf("add feed expects 2 arguments, the name and url")
    }
    ctx := context.Background()
    name := cmd.Args[0]
	url := cmd.Args[1]
	username := s.Cfg.CurrentUserName
	user, err := s.Db.GetUser(ctx, username)
	if err != nil {
		return err
	}
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	}
    feed, err := s.Db.CreateFeed(ctx, feedParams)
    if err != nil {
        return err
    }
    fmt.Print(feed)
	    followParams := database.CreateFeedFollowParams{
        ID:        uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        UserID:    user.ID,
        FeedID:    feed.ID,
    }
    _, err = s.Db.CreateFeedFollow(ctx, followParams)
    if err != nil {
        return fmt.Errorf("could not create feed follow: %w", err)
    }
    return nil
}