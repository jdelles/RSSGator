// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: getFeedByURL.sql

package database

import (
	"context"
)

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at 
FROM feeds
WHERE url = $1
`

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}
