package upload

// ArtifactUpload is provided from a client request uploading a file
type ArtifactUpload struct {
	Path          string // Path local to build
	FormKey       string // Form key safe string used as key
	Label         string // "Arbitrary Title"
	PostProcessor string // "cobertura:xml" | "codestyle" | etc seek docs
	Archived      bool   // Whether we are recieving a zip archive of files
}
