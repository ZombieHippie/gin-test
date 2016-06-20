
import { Repository } from "../repo/repo.model"
import { Artifact } from "../artifact/artifact.model"

interface Summary {
  ID?:          number // Assigned by gorm
  Repository:   Repository // Assign this yourself
  RepositoryID: string
  BranchID:     string // "feature/no-more-bugs"
  BuildID:      number // 2
  Commit:       string // Commit hash

  Author?:  string
  Message?: string
  Success:  boolean
  Created:  Date // Date Summary was recorded
}

export { Summary }
