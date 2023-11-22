// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: feed_follows.sql

package database

import (
	"context"
	"database/sql"
)

const createFeedFollow = `-- name: CreateFeedFollow :execresult
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id) VALUES (UUID(), NOW(), ?, ?, ?)
`

type CreateFeedFollowParams struct {
	UpdatedAt sql.NullTime
	UserID    string
	FeedID    string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createFeedFollow, arg.UpdatedAt, arg.UserID, arg.FeedID)
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :execresult
delete from feed_follows where id = ? and user_id = ?
`

type DeleteFeedFollowParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
}

const getFeedFollows = `-- name: GetFeedFollows :many
select id, created_at, updated_at, user_id, feed_id from feed_follows where user_id = ?
`

func (q *Queries) GetFeedFollows(ctx context.Context, userID string) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}