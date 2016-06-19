package artifact

import (
	"github.com/ZombieHippie/test-gin/server/src/shared"
)

// PostProcess uses the PostProcessor attribute of the Artifact to assign
// its status and data properties.
func (art *Artifact) PostProcess() error {
	art.Status = shared.StatusWarn
	if art.PostProcessor == "cobertura" {
		art.Status = shared.StatusSuccess
	}
	// TODO: Implement actual handling of these different post processors

	art.Data = `{ "Message": "Hello" }`
	return nil
}
