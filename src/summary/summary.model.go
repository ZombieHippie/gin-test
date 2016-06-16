package summary

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/jinzhu/gorm"
)

// Summary is a common object for reports to point at
type Summary struct {
	gorm.Model
	Repository    repo.Repository
	PullRequestID int64
	Commit        string
	Message       string
	Author        string
	Success       bool
	Created       int64
}
