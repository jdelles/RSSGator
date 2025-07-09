package handlers

import (
	"RSSGator/commands"
	"RSSGator/functions"
	"context"
	"fmt"
	"os"
	"time"
)

func HandlerAgg(s *commands.State, cmd commands.Command) error {
    if len(cmd.Args) == 0 {
        return fmt.Errorf("agg expects a single argument: time_between_reqs (e.g., 1s, 1m, 1h)")
    }

    timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
    if err != nil {
        return fmt.Errorf("invalid duration: %w", err)
    }

    fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
    ticker := time.NewTicker(timeBetweenRequests)
    defer ticker.Stop()

    ctx := context.Background()

    for {
        fmt.Println("Scraping feeds...")
        if err := functions.ScrapeFeeds(ctx, s.Db); err != nil {
            fmt.Fprintf(os.Stderr, "Error scraping feeds: %v\n", err)
        }
        <-ticker.C
    }
}
