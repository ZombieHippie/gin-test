package report

import (
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
)

// TestReport is created if a test report exists
type TestReport struct {
	gorm.Model
	Summary      summary.Summary
	FileContents string
	Label        string
	Passed       int64
	Failed       int64
}

// TestError stores individual test failure error information from report
type TestError struct {
	gorm.Model
	TestReport TestReport
	Level      int64
	Position   int64
	Path       string
	Note       string
}
