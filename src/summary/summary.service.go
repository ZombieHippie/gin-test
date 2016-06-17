package summary

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/jinzhu/gorm"
)

// GetSummariesByBranch retrieves the the Summary by repo and pull request
func GetSummariesByBranch(db *gorm.DB, repoID string, branchID string) []Summary {
	var repo repo.Repository
	db.First(&repo, repoID)

	var sums = make([]Summary, 16)
	// Latest summary
	db.Where(&Summary{
		BranchID:   branchID,
		Repository: repo,
	}).Order("summary_id").Find(&sums)
	return sums
}

// GetAllSummaries retrieves all the Summaries
func GetAllSummaries(db *gorm.DB) ([]Summary, int) {
	var sums = make([]Summary, 16)
	var count int
	db.Find(&sums).Count(&count)
	return sums, count
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
