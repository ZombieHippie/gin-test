package upload

// ArtifactUpload is provided from a client request uploading a file
type ArtifactUpload struct {
	Path          string // Path local to build
	Label         string // "Arbitrary Title"
	PostProcessor string // "cobertura:xml" | "codestyle" | etc seek docs
}
