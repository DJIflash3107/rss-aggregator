// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	Name          string
	Url           string
	UserID        string
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	UserID    string
	FeedID    string
}

type Post struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	Title       string
	Description sql.NullString
	PublishedAt time.Time
	Url         string
	FeedID      string
}

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Name      string
	ApiKey    string
}
