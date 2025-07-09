package handlers

import (
	"RSSGator/commands"
	"context"
	"fmt"
)

func HandlerFeeds(s *commands.State, _ commands.Command) error {
	ctx := context.Background()
	feeds, err := s.Db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		if i != 0 {
			println("##############################")
		}
		user, err := s.Db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}
	return nil
}
