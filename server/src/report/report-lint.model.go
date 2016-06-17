package report

import (
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/jinzhu/gorm"
)

// LintReport stores information about linting results
type LintReport struct {
	gorm.Model
	Summary  summary.Summary
	Label    string
	Errors   int64
	Warnings int64
}

// LintError stores individual linting error information from report
type LintError struct {
	gorm.Model
	LintReport LintReport
	Level      int64
	Position   int64
	Path       string
	Note       string
}
