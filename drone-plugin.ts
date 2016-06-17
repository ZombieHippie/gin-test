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
  auth: string // "authy",
  files: VFile[]
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
      return createArtifact(fullpath, file.label, file.filename, !!file.binary)
    })

    const summary: Summary = {
      Branch: p.build.branch,
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

    UploadSummary(host, 'authy', summary, (err, summary) => {
      console.log(err, summary)
    })

    vargs

  } else {
    console.log("Parameter missing: Server host");
    process.exit(1)
  }  
}

/*
const host = 'localhost:8080'



const PromiseSftp = require('promise-sftp');

const path = require('path');
const shelljs = require('shelljs');

const do_upload = function (workspace, vargs) {
  if (vargs.host) {

    var sftp = new PromiseSftp();
    vargs.destination_path || (vargs.destination_path = '/');

    sftp.connect({
      host: vargs.host,
      port: vargs.port,
      username: vargs.username,
      password: vargs.password,
      privateKey: workspace.keys && workspace.keys['private']
    }).then(function (greetings) {
      console.log('Connection successful. ' + (greetings || ''));
     
      return [].concat.apply([], vargs.files.map((f) => { return shelljs.ls(workspace.path + '/' + f); }));
    }).each(function(file) {
      var basename = path.basename(file);

      console.log('Uploading ' + file + ' as ' + basename + ' into ' + vargs.destination_path);
      return sftp.put(file, path.join(vargs.destination_path, basename))
    }).then(function() {

      console.log('Upload successful');
    }).catch(function(err) {

      console.log('An error happened: ' + err);
      process.exit(2)
    }).then(function() {

      sftp.logout();
    });
  } else {
    console.log("Parameter missing: SFTP server host");
    process.exit(1)
  }
}

plugin.parse().then((params) => {

  // gets build and repository information for
  // the current running build
  const build = params.build;
  const repo  = params.repo;
  const workspace = params.workspace;

  // gets plugin-specific parameters defined in
  // the .drone.yml file
  const vargs = params.vargs;

  vargs.username      || (vargs.username = '');
  vargs.files         || (vargs.files = []);

  do_upload(workspace, vargs);
});

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
  FileContents: readFileSync('./.editorconfig', 'utf8'),
  IsBinary: false,
  FileName: "unittests",
  Data: `{
    "pass": 12,
    "fail": 0,
    "error": 0

  }`,
  Label: 'unit',
  Passed: 12,
  Failed: 0,
}

*/

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

