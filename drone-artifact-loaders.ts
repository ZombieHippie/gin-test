import { Artifact } from "./src/artifact/artifact.model"

// loaders for applying additional information based on label
const Loaders: { [label: string]: (art: Artifact) => void } = {
  // each one can populate the Data, Pass, and Fail attributes of the Artifact
}

export { Loaders }
