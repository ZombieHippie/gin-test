
import { Repository } from "../repo/repo.model"
import { ArtifactUpload } from "../artifact/artifact-upload.model"

interface SummaryUpload {
  Repository: Repository // Assign this yourself
  BuildID: number // 2
  BranchID: string // "feature/no-more-bugs"
  Artifacts: ArtifactUpload[] // Files you want to attach
  Commit: string // Commit hash
  Success: boolean
  Created: Date // Date Summary was recorded
  Message?: string
  Author?: string
}

export { SummaryUpload }
