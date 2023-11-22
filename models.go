package main

import (
	"time"

	"github.com/DJIflash3107/rss-aggregator/internal/database"
)

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Name      string     `json:"name"`
	ApiKey    string     `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	var updated_at *time.Time
	if dbUser.UpdatedAt.Valid {
		updated_at = &dbUser.UpdatedAt.Time
	}
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: updated_at,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	UserID    string     `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	var updated_at *time.Time
	if dbFeed.UpdatedAt.Valid {
		updated_at = &dbFeed.UpdatedAt.Time
	}
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: updated_at,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UserID    string     `json:"user_id"`
	FeedID    string     `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollows database.FeedFollow) FeedFollow {
	var updated_at *time.Time
	if dbFeedFollows.UpdatedAt.Valid {
		updated_at = &dbFeedFollows.UpdatedAt.Time
	}
	return FeedFollow{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: updated_at,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.FeedID,
	}
}
func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	PublishedAt time.Time  `json:"published_at"`
	Url         string     `json:"url"`
	FeedID      string     `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	var updated_at *time.Time
	if dbPost.UpdatedAt.Valid {
		updated_at = &dbPost.UpdatedAt.Time
	}
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   updated_at,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}
