-- name: CreateFeedFollow :execresult
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id) VALUES (UUID(), NOW(), ?, ?, ?);

-- name: GetFeedFollows :many
select * from feed_follows where user_id = ?;

-- name: DeleteFeedFollow :execresult
delete from feed_follows where id = ? and user_id = ?;