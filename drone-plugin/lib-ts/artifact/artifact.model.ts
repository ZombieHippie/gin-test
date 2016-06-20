
import { Summary } from "../summary/summary.model"

interface Artifact {
  Summary:       Summary
	SummaryID:     number
	LocalPath:     string // Path located from build
	Label:         string // "Arbitrary Title"
	PostProcessor: string // "cobertura"
  Path:          string // /data/{repo}/{branch}/{build number}/{file path}
	Data:          string // Some JSON formatted data?
	Status:        string // "pass", "fail", "error", "warn"
}

export { Artifact }
