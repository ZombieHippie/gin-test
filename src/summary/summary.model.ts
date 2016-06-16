
import { Repository } from "../repo/repo.model"
import { Artifact } from "../artifact/artifact.model"

interface Summary {
  ID?: number // Assigned by gorm
  Repository: Repository // Assign this yourself
  PullRequestID: number // "user:ZombieHippie,org:DryClean"
  Artifacts: Artifact[] // Files you want to attach
  Commit: string // Commit hash
  Success: boolean
  Created: Date // Date Summary was recorded
  Message?: string
  Author?: string
}

export { Summary }

