package artifact

import (
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
)

// Artifact is created for each piece generated in a summary
type Artifact struct {
	gorm.Model
	Summary      summary.Summary
	FileContents string
	Data         string // Some JSON formatted data?
	Label        string
	Passed       int64
	Failed       int64
}
