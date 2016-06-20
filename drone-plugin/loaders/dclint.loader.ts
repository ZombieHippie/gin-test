/// <reference path="../typings/index.d.ts" />

import { ArtifactUpload } from '../lib-ts/upload/artifact-upload.model'
import { LoaderFactory, Loader } from './loader.interface'
import { PostComment } from '../lib-ts/github/post-comment.function'

import fs = require('fs')

export function DCLintLoader(query: {[key: string]: string }): Loader {
  console.log("DCLintLoader recieved:", query)

  return function (artifact: ArtifactUpload): Error {
    if (query["format"] == "tslint-prose") {
      let file = fs.readFileSync(artifact.Path, 'utf8')
      let lines = file.split(/\s*\n\s*/g)

      // src/app/app.component.ts[66, 1]: exceeds maximum line length of 50

      let TSLintProseRE = /^([^[]+)\[(\d+), \d+\]: (.+)$/

      lines.forEach((line) => {
        let res = TSLintProseRE.exec(line)
        if (res != null) {
          let [_, filePath, pos, comment] = res
          console.log("TSLINT:", filePath, pos, comment)
          // PostComment(filePath, parseInt(pos), comment, 1)
        }
      })
    } else {
      return new Error("DCLint can only interpret tslint-prose at the moment.")
    }
    return null
  }
} 
