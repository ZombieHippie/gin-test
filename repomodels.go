package main

import "github.com/jinzhu/gorm"
import "github.com/ZombieHippie/test-gin/src/summary"

// GithubCommitComment stores a reference to a commit that has been made
type GithubCommitComment struct {
	gorm.Model
	Summary         summary.Summary
	GithubCommentID int64
	Label           string
	Level           int64
	Position        int64
	Path            string
}
