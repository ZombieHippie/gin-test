package summary

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/jinzhu/gorm"
	"log"
)

// GetSummariesByPullRequest retrieves the the Summary by repo and pull request
func GetSummariesByPullRequest(db *gorm.DB, repoID string, pr shared.PullRequest) []Summary {
	var repo repo.Repository
	db.First(&repo, repoID)

	var sums = make([]Summary, 16)
	// Latest summary
	db.Where(&Summary{
		PullRequestID: pr,
		Repository:    repo,
	}).Order("summary_id").Find(&sums)
	return sums
}

// GetAllSummaries retrieves all the Summaries
func GetAllSummaries(db *gorm.DB) []Summary {
	var sums = make([]Summary, 16)
	var count int
	db.Find(&sums).Count(count)
	log.Println(sums, count)
	return sums
}

// CreateSummary inserts a Summary into the database
func CreateSummary(db *gorm.DB, sum Summary) (bool, Summary) {

	created := db.NewRecord(sum) // => returns `true` as primary key is blank

	if !created {
		return false, sum
	}

	db.Create(&sum)

	return true, sum
}
