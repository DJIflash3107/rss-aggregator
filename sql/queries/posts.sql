-- name: CreatePost :execresult
INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id) VALUES (UUID(), NOW(), ?, ?, ?, ?, ?, ?);

-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = ?
ORDER BY posts.published_at DESC
LIMIT ?;