-- name: CreateFeed :execresult
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) VALUES (UUID(), NOW(), ?, ?, ?, ?);

-- name: GetFeeds :many
select * from feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at IS NULL, last_fetched_at ASC
LIMIT ?;

-- name: MarkFeedAsFetched :execresult
update feeds set last_fetched_at = NOW(), updated_at = NOW() where id = ?;