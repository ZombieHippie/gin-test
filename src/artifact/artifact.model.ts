
interface Artifact {
  FileContents?: string
  IsBinary: boolean
  FileName: string
  Data?: string
  Label: string
  Passed: number
  Failed: number
}

export { Artifact }
