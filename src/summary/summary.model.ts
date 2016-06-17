
import { Repository } from "../repo/repo.model"
import { Artifact } from "../artifact/artifact.model"

interface Summary {
  ID?: number // Assigned by gorm
  Repository: Repository // Assign this yourself
  BuildID: number // 2
  PullRequestID: number // 5
  BranchID: string // "feature/no-more-bugs"
  Artifacts: Artifact[] // Files you want to attach
  Commit: string // Commit hash
  Success: boolean
  Created: Date // Date Summary was recorded
  Message?: string
  Author?: string
}

export { Summary }

