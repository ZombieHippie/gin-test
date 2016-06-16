package summary

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/report"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/jinzhu/gorm"
)

// GetSummariesByPullRequest retrieves the the Summary by repo and pull request
func GetSummariesByPullRequest(db *gorm.DB, repoID repo.RepoID, pr shared.PullRequest) (sums []Summary) {
	var repo repo.Repository
	db.First(&repo, repoID)
	// Latest summary
	db.Where(&Summary{
		PullRequestID: pr,
		Repository:    repo,
	}).Order("summary_id").Find(&sums)
	return
}

// GetAllSummaries retrieves all the Summaries
func GetAllSummaries(db *gorm.DB) (sums []Summary) {
	db.Order("summary_id").Find(&sums)
	return
}

// CreateSummary inserts a Summary into the database
func CreateSummary(db *gorm.DB, sum Summary) (bool, Summary) {

	created := db.NewRecord(sum) // => returns `true` as primary key is blank

	if !created {
		return false
	}

	db.Create(&sum)

	return true, sum
}
