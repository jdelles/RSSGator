package handlers


import (
	"context"
	"fmt"
	"RSSGator/commands"
	"RSSGator/functions"
)

func HandlerAgg(s *commands.State, cmd commands.Command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	feed, err := functions.FetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Print(feed)
	return nil
}