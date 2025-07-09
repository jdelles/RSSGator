package handlers

import (
	"context"
	"fmt"

	"RSSGator/commands"
	"RSSGator/internal/database"
)

func HandlerUnfollow(s *commands.State, cmd commands.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("unfollow command requires exactly 1 argument: url")
	}
	feedURL := cmd.Args[0]
	ctx := context.Background()

	params := database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		Url: feedURL,
	}

	err := s.Db.DeleteFeedFollowByUserAndURL(ctx, params)
	if err != nil {
		return fmt.Errorf("could not delete feed with url: %s", err)
	}

	fmt.Printf("User %s is no longer following feed %s", user.Name, feedURL)
	return nil
}