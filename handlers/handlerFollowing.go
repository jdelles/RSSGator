package handlers

import (
	"context"
	"fmt"

	"RSSGator/commands"
	"RSSGator/internal/database"
)

func HandlerFollowing(s *commands.State, _ commands.Command, user database.User) error {
	ctx := context.Background()

	follows, err := s.Db.GetFeedFollowsByUser(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}

// Add a following command. It should print all the names of the feeds the current user is following.
