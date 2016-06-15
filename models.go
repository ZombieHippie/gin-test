package main

// Article is a post in the database
type Article struct {
	ID      int64 `db:"article_id"`
	Created int64
	Title   string
	Content string
}
