package upload

import (
	"github.com/ZombieHippie/test-gin/server/src/repo"
	"time"
)

// SummaryUpload is a common object for reports to point at
type SummaryUpload struct {
	Repository repo.Repository
	BranchID   string
	BuildID    int
	Artifacts  []ArtifactUpload
	Commit     string

	Message string
	Author  string
	Success bool
	Created time.Time
}
