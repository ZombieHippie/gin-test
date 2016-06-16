package summary

import (
	"github.com/ZombieHippie/test-gin/src/artifact"
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/jinzhu/gorm"
	"time"
)

// Summary is a common object for reports to point at
type Summary struct {
	gorm.Model
	Repository    repo.Repository
	PullRequestID shared.PullRequest
	Artifacts     []artifact.Artifact
	Commit        shared.Commit
	Message       string
	Author        string
	Success       bool
	Created       time.Time
}

// TableName sets Summary's table name to be `summaries``
func (Summary) TableName() string {
	return "summaries"
}
