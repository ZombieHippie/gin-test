package artifact

import (
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/jinzhu/gorm"
)

// Artifact is created for each piece generated in a summary
type Artifact struct {
	gorm.Model
	Summary       summary.Summary
	LocalPath     string // Path located from build
	Label         string // "Arbitrary Title"
	PostProcessor string // "cobertura"
	Path          string // Path located on server
	Data          string // Some JSON formatted data?
	Status        string // "pass", "fail", "error", "warn"
}
