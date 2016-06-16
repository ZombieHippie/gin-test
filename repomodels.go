package main

// Repository is parent to almost every other model
type Repository struct {
	ID     string // `db:"repo_id"`
	ACL    string
	Owner  string
	Name   string
	Active bool
}

// Summary is a common object for reports to point at
type Summary struct {
	SummaryID     int64  `db:"summary_id, primarykey, autoincrement"`
	RepoID        string `db:"repo_id"`
	PullRequestID int64  `db:"pull_request_id"`
	Commit        string
	Message       string
	Author        string
	Success       bool
	Created       int64
}

// CoverageReport is used to store results of test coverage
type CoverageReport struct {
	SummaryID    int64  `db:"summary_id"`
	FileContents string `db:"file_contents"`
	Label        string
	Coverage     float64
}

// TestReport is created if a test report exists
type TestReport struct {
	SummaryID    int64  `db:"summary_id"`
	FileContents string `db:"file_contents"`
	Label        string
	Passed       int64
	Failed       int64
}

// LintReport stores information about linting results
type LintReport struct {
	SummaryID int64 `db:"summary_id"`
	Label     string
	Errors    int64
	Warnings  int64
}

// LintError stores individual linting error information from report
type LintError struct {
	SummaryID int64 `db:"summary_id"`
	Label     string
	Level     int64
	Position  int64
	Path      string
	Note      string
}

// GithubCommitComment stores a reference to a commit that has been made
type GithubCommitComment struct {
	SummaryID       int64 `db:"summary_id"`
	GithubCommentID int64
	Label           string
	Level           int64
	Position        int64
	Path            string
}
