package report

import (
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
)

// CoverageReport is used to store results of test coverage
type CoverageReport struct {
	gorm.Model
	Summary      summary.Summary
	FileContents string
	Label        string
	Coverage     float64
}
