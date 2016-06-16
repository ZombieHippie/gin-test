package artifact

import (
	"github.com/jinzhu/gorm"
)

// Artifact is created for each piece generated in a summary
type Artifact struct {
	gorm.Model
	FileContents string
	Data         string // Some JSON formatted data?
	Label        string
	Passed       int64
	Failed       int64
}
