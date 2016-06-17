/// <reference path="./typings/index.d.ts" />
import { UploadSummary } from "./lib-ts/app/upload-summary"
import { Summary } from "./lib-ts/summary/summary.model"
import { Artifact } from "./lib-ts/artifact/artifact.model"

const host = 'localhost:8080'

import { readFileSync } from "fs"


const lintReport: Artifact = {
  FileContents: readFileSync('./large-image-test.jpg', 'base64'),
  IsBinary: true,
  FileName: "lint.jpg",
  Data: `{
    "error": 0,
    "warning": 4
  }`,
  Label: 'lint',
  Passed: 1, // has no significance
  Failed: 0,
}
const testReport: Artifact = {
  FileContents: readFileSync('./drone-plugin.ts', 'utf8'),
  IsBinary: false,
  FileName: "unit",
  Data: `{
    "pass": 12,
    "fail": 0,
    "error": 0

  }`,
  Label: 'unit',
  Passed: 12,
  Failed: 0,
}

const arts: Artifact[] = [
  lintReport,
  testReport
]


for (var i = 0; i < 5; i++) {
  const summary: Summary = {
    BranchID: "feature/hello",
    BuildID: i,
    Commit: "962c4b831f447bccd8ab4185a4898d41833d91d3",
    Author: "Cole R Lawrence <colelawr@gmail.com>",
    Message: "Fix all golang compilation errors",
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
