
interface ArtifactUpload {
  Path:          string // Local path for metadata
  Label:         string
  PostProcessor: string
  FormKey?:      string // Form key safe string
  Archived?:     boolean
}

export { ArtifactUpload }
