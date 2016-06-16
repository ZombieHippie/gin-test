package repo

import (
	"github.com/ZombieHippie/test-gin/src/report"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
)

// EnsureRepository inserts a Repository into the database if it doesn't exist and returns that Repository
func EnsureRepository(db *gorm.DB, repo Repository) Repository {
	db.FirstOrCreate(&repo, repo)
	return repo
}
