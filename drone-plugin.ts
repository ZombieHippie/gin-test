/// <reference path="./typings/index.d.ts" />
/// <reference path="./drone-node.d.ts" />

const Drone = require('drone-node')
import shelljs = require('shelljs')
const plugin = new Drone.Plugin()

import path = require('path')
import { readFileSync } from "fs"

import { UploadSummary } from "./src/app/upload-summary"
import { Summary } from "./src/summary/summary.model"
import { Artifact } from "./src/artifact/artifact.model"
import { Loaders } from "./drone-artifact-loaders"


interface Vargs {
  host: string // "art.company.com:8080",
  files: VFile[]
  auth?: string // "authy",
  loadersDir?: string // ".config/drone-loaders"
}

interface VFile {
  path: string // "stats.json", "coverage/Phantom*/index.html"
  label: string // "webpack", "coverage"
  filename?: string // "webpack.stats.json", "coverage.html"
  binary?: boolean
}

function postSummary(p: DroneParams, vargs: Vargs) {
  if (vargs.host) {
    const arts = vargs.files.map<Artifact>((file: VFile) => {
      let fullpath = shelljs.ls(path.join(p.workspace.path, file.path))[0]
      file.filename = file.filename === "" ? file.label : file.filename
      return createArtifact(fullpath, file.label, file.filename, !!file.binary)
    })

    const summary: Summary = {
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
      console.log(resp.Summary.Artifacts.map((art) => art.FileContents).join("\n"))
    })
  } else {
    console.log("Parameter missing: Server host")
    process.exit(1)
  }  
}

function createArtifact(filepath: string, label: string, filename: string, isBinary: boolean): Artifact {
  let art:Artifact =  {
    FileContents: readFileSync(filepath, isBinary ? 'base64' : 'utf8'),
    IsBinary: isBinary,
    FileName: filename,
    Data: '',
    Label: label,
    Passed: 1,
    Failed: 0,
  }
  if (typeof Loaders[label] === 'function') {
    Loaders[label](art)
  }
  return art
}

plugin.parse().then((params: DroneParams) => {
  // gets plugin-specific parameters defined in
  // the .drone.yml file
  const vargs = params.vargs as Vargs

  postSummary(params, vargs)
})
