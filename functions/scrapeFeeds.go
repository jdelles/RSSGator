package functions

import (
	"RSSGator/internal/database"

	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func parsePublishedAt(dateStr string) (sql.NullTime, error) {
    layouts := []string{
        time.RFC1123Z,
        time.RFC1123,
        time.RFC822Z,
        time.RFC822,
        time.RFC3339,
        "Mon, 02 Jan 2006 15:04:05 MST", 
    }
    for _, layout := range layouts {
        if t, err := time.Parse(layout, dateStr); err == nil {
            return sql.NullTime{Time: t, Valid: true}, nil
        }
    }
    return sql.NullTime{Valid: false}, fmt.Errorf("could not parse published_at: %s", dateStr)
}

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {
	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("could not get next feed to fetch: %w", err)
	}

	if err := db.MarkFeedFetched(ctx, feed.ID); err != nil {
		return fmt.Errorf("could not mark feed as fetched: %w", err)
	}

	rssFeed, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

    now := time.Now()
    for _, item := range rssFeed.Channel.Item {
        publishedAt, _ := parsePublishedAt(strings.TrimSpace(item.PubDate))

        params := database.CreatePostParams{
            ID:          uuid.New(),
            CreatedAt:   now,
            UpdatedAt:   now,
            Title:       item.Title,
            Url:         item.Link,
            Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
            PublishedAt: publishedAt,
            FeedID:      feed.ID,
        }

        _, err := db.CreatePost(ctx, params)
        if err != nil {
            if strings.Contains(err.Error(), "duplicate key value") {
                continue
            }
            fmt.Printf("error saving post '%s': %v", item.Title, err)
        }
    }

	return nil
}
