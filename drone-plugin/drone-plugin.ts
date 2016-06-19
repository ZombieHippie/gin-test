/// <reference path="./typings/index.d.ts" />
/// <reference path="./drone-node.d.ts" />

const Drone = require('drone-node')
import shelljs = require('shelljs')
const plugin = new Drone.Plugin()

import path = require('path')
import { readFileSync } from "fs"

import { UploadSummary } from "./lib-ts/app/upload-summary"
import { SummaryUpload } from "./lib-ts/upload/summary-upload.model"
import { ArtifactUpload } from "./lib-ts/upload/artifact-upload.model"


interface Vargs {
  host:         string // "art.company.com:8080",
  files:        VFile[]
  auth?:        string // "authy"
}

interface VFile {
  path:           string // "stats.json", "coverage/Phantom*/index.html"
  label:          string // "webpack", "coverage"
  postprocessor?: string // "cobertura", "codestyle", etc
}

function postSummary(p: DroneParams, vargs: Vargs) {
  if (vargs.host) {
    const arts = vargs.files.map<ArtifactUpload>((file: VFile) => {
      let fullpath = shelljs.ls(path.join(p.workspace.path, file.path))[0]
      return {
        Label:          file.label,
        LocalPath:      file.path,
        Path:           fullpath,
        PostProcessor:  file.postprocessor,
      }
    })

    const summary: SummaryUpload = {
      BranchID: p.build.branch,
      BuildID: p.build.number,
      Commit: p.build.commit,
      Author: p.build.author,
      Message: p.build.message,
      Artifacts: arts,
      Success: p.build.status === "success",
      Created: new Date(p.build.finished_at),
      Repository: {
        ID: "ZombieHippie/test-gin",
        ACL: "user:ZombieHippie",
        Active: true
      }
    }

    UploadSummary(vargs.host, vargs.auth, summary, (err, resp) => {
      if (err) {
        console.error("Error occurred while posting artifacts!", err)
        // Don't mark the build as failing just because we can't post.
        process.exit(0)
      }
      console.log("Uploaded files:")
      console.log(resp.Artifacts.map((art) => `  ${art.Label} (${art.Status}): ${art.Data}`).join("\n"))
    })
  } else {
    console.log("Parameter missing: Server host")
    process.exit(1)
  }  
}

import { inspect } from "util"

plugin.parse().then((params: DroneParams) => {
  // gets plugin-specific parameters defined in
  // the .drone.yml file
  const vargs = params.vargs as Vargs

  console.log("params: DroneParams = ", inspect(params, false, 8, true))

  postSummary(params, vargs)
})
