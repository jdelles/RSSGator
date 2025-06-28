-- name: GetFeedFollowsByUser :many
SELECT 
    ff.*,
    u.name AS user_name,
    f.name AS feed_name
FROM 
    feed_follows ff
JOIN 
    users u on ff.user_id = u.id
JOIN 
    feeds f on ff.feed_id = f.id
WHERE 
    ff.user_id = $1;