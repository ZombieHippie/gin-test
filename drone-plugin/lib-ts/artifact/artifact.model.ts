
import { Summary } from "../summary/summary.model"

interface Artifact {
  Summary:       Summary
  Path:          string // /data/{repo}/{branch}/{build number}/{file path}
  FileType:      string // "xml", "json", "html", "directory", etc
  PostProcessor: string // "cobertura"
  Data?:         string // cobertura-output-data
  Label:         string // "Arbitrary title"
  Status:        string // "pass", "fail", "error", "warn"
}

export { Artifact }
