package handlers

import (
	"RSSGator/commands"
	"RSSGator/internal/database"

	"context"
	"fmt"
	"strconv"

)

func HandlerBrowse(s *commands.State, cmd commands.Command, user database.User) error {
    ctx := context.Background()
    limit := 2

    if len(cmd.Args) > 0 {
        if n, err := strconv.Atoi(cmd.Args[0]); err == nil && n > 0 {
            limit = n
        }
    }

    posts, err := s.Db.GetPostsForUser(ctx, database.GetPostsForUserParams{
        UserID: user.ID,
        Limit:  int32(limit),
    })
    if err != nil {
        return fmt.Errorf("could not get posts: %w", err)
    }

    for _, post := range posts {
        fmt.Printf("Title: %s\nURL: %s\nPublished: %v\nDescription: %s\n---\n",
            post.Title, post.Url, post.PublishedAt.Time, post.Description.String)
    }

    return nil
}
