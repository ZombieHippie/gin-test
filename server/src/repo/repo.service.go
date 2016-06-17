package repo

import (
	"github.com/jinzhu/gorm"
)

// EnsureRepository inserts a Repository into the database if it doesn't exist and returns that Repository
func EnsureRepository(db *gorm.DB, given Repository) (repo Repository) {
	db.FirstOrCreate(&repo, given)
	return
}
