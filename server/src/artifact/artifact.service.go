package artifact

import (
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/jinzhu/gorm"
)

// GetAllArtifacts retrieves all the Artifacts
func GetAllArtifacts(db *gorm.DB) ([]Artifact, int) {
	var arts = make([]Artifact, 16)
	var count int
	db.Find(&arts).Count(&count)
	return arts, count
}

// GetArtifacts retrieves all the Artifacts created from buildID
func GetArtifacts(db *gorm.DB, sum *summary.Summary) ([]Artifact, int) {
	var arts = make([]Artifact, 16)
	var count int
	db.Find(&arts, Artifact{
		Summary: *sum,
	}).Count(&count)
	return arts, count
}
