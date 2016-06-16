package report

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
)

// GetTest retrieves the TestReport by ID
func GetTest(db *gorm.DB, testID shared.ModelID) (result shared.DatabaseResult, report TestReport) {
	db.First(&report, testID)
	return
}

// GetTestByPullRequest retrieves the TestReport by ID
func GetTestByPullRequest(db *gorm.DB, repoID shared.ModelID, pr shared.PullRequest) (result shared.DatabaseResult, report TestReport) {
	sum := summary.GetSummariesByPullRequest(db, repoID, pr)[0]
	db.Where(&TestReport{Summary: sum}).First(&report)
	return
}
