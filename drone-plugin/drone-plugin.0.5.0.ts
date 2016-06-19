/// <reference path="./typings/index.d.ts" />
/// <reference path="./drone-node.d.ts" />

import shelljs = require('shelljs')

import path = require('path')
import { readFileSync } from "fs"

import { UploadSummary } from "./lib-ts/app/upload-summary"
import { SummaryUpload } from "./lib-ts/upload/summary-upload.model"
import { ArtifactUpload } from "./lib-ts/upload/artifact-upload.model"

import { ENV, PLUGIN_ENV, VFile } from "./drone-parser"

function postSummary(vargs: PLUGIN_ENV) {
  if (vargs.PLUGIN_HOST) {
    const arts = vargs.PLUGIN_FILES.map<ArtifactUpload>((file: VFile) => {
      // cwd is at the route of the project
      let fullpath = shelljs.ls(file.path)[0]
      return {
        Label:          file.label,
        LocalPath:      file.path,
        Path:           fullpath,
        PostProcessor:  file.postprocessor,
      }
    })

    const summary: SummaryUpload = {
      BranchID: vargs.DRONE_BRANCH,
      BuildID:  vargs.DRONE_BUILD_NUMBER,
      Commit:   vargs.DRONE_COMMIT,
      Author: 	vargs.DRONE_COMMIT_AUTHOR,
      Message:  vargs.DRONE_COMMIT_MESSAGE,
      Artifacts: arts,
      Success: vargs.DRONE_BUILD_STATUS === "success",
      Created: new Date(vargs.DRONE_BUILD_CREATED),
      Repository: {
        ID:   vargs.DRONE_REPO,
        ACL:  "user:ZombieHippie",
        Active: true
      }
    }

    UploadSummary(vargs.PLUGIN_HOST, vargs.PLUGIN_AUTH, summary, (err, resp) => {
      if (err) {
        console.error("Error occurred while posting artifacts!", err)
        // Don't mark the build as failing just because we can't post.
        process.exit(0)
      }
      console.log("Uploaded files:")
      try {
        console.log(resp.Artifacts.map((art) => `  ${art.Label} (${art.Status}): ${art.Data}`).join("\n"))
      } catch (err) {
        console.log("Error:", resp)
      }
    })
  } else {
    console.log("Parameter missing: Server host")
    process.exit(1)
  }  
}

import { inspect } from "util"

// gets plugin-specific parameters defined in
// the .drone.yml file
console.log("params: DroneParams = ", inspect(ENV, false, 8, true))
postSummary(ENV)
