-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    last_fetched_at = NOW(),
    updated_at = NOW()
WHERE 
    ID = $1;