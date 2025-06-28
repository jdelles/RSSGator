package handlers

import (
	"context"
	"fmt"

	"RSSGator/commands"
)

func HandlerFollowing(s *commands.State, _ commands.Command) error {
	ctx := context.Background()
	user, err := s.Db.GetUser(ctx, s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

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