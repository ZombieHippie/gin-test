
interface ArtifactUpload {
  Path:          string // Local path for metadata
  Label:         string
  PostProcessor: string
  Archived?:      boolean
}

export { ArtifactUpload }
