package handlers

import (
	"context"
	"fmt"
	"time"

	"RSSGator/commands"
	"RSSGator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command requires exactly 1 argument: url")
	}
	feedURL := cmd.Args[0]
	ctx := context.Background()

	user, err := s.Db.GetUser(ctx, s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("count not find current user: %s", err)
	}

	feed, err := s.Db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("count not find feed with url: %s", err)
	}

	params := database.CreateFeedFollowParams {
		ID:        uuid.New(), 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.Db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("count not create feed follow: %s", err)
	}

	fmt.Printf("User %s is now following feed %s", follow.UserName, follow.FeedName)
	return nil
}