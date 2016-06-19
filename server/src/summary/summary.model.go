package summary

import (
	"github.com/ZombieHippie/test-gin/server/src/repo"
	"github.com/jinzhu/gorm"
	"time"
)

// Summary is a common object for reports to point at
type Summary struct {
	gorm.Model
	Repository   repo.Repository `gorm:"ForeignKey:RepositoryID"`
	RepositoryID string
	BranchID     string
	BuildID      int
	Commit       string

	Message string
	Author  string
	Success bool
	Created time.Time
}
