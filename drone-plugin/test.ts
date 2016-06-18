/// <reference path="./typings/index.d.ts" />
import { UploadSummary } from "./lib-ts/app/upload-summary"
import { SummaryUpload } from "./lib-ts/summary/summary-upload.model"
import { ArtifactUpload } from "./lib-ts/artifact/artifact-upload.model"

const host = 'localhost:8080'

import { readFileSync } from "fs"


const coverageReport: ArtifactUpload = {
  Path:           './mocks/coverage-with-data.xml',
  Label:          'Coverage',
  PostProcessor:  'cobertura'
}
const lintReport: ArtifactUpload = {
  Path:           './mocks/large-image-test.jpg',
  Label:          'Surfing',
  PostProcessor:  'image'
}
const testReport: ArtifactUpload = {
  Path:           './drone-plugin.ts',
  Label:          'Unit Tests',
  PostProcessor:  'junit',
}

const arts: ArtifactUpload[] = [
  lintReport,
  testReport,
  coverageReport
]


for (var i = 0; i < 5; i++) {
  const summary: SummaryUpload = {
    BranchID: "feature/hello",
    BuildID: i,
    Commit: i + "62c4b831f447bccd8ab4185a4898d41833d91d3",
    Author: "Cole R Lawrence <colelawr@gmail.com>",
    Message: i + " Fix all golang compilation errors",
    Artifacts: arts,
    Success: true,
    Created: new Date(),
    Repository: {
      ID: "ZombieHippie/test-gin",
      ACL: "user:ZombieHippie",
      Active: true
    }
  }

  UploadSummary(host, 'authy', summary, (err, summary) => {
    console.log(err, summary)
  })  
}
