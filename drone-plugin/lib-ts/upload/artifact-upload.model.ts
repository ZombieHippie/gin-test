
interface ArtifactUpload {
  Path:          string // Local path for metadata
  FullPath?:     string // Full path including workspace for uploading
  Label:         string
  PostProcessor: string
}

export { ArtifactUpload }
