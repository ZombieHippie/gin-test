/// <reference path="../typings/index.d.ts" />


import { ArtifactUpload } from '../lib-ts/upload/artifact-upload.model'
import { Loader } from './loader.interface'

import querystring = require('querystring')

import { DCLintLoader } from './dclint.loader'

const loaders: {[key: string]: (query: {[key: string]: string }) => Loader} = {
  "dc-lint": DCLintLoader
}

export function GetLoader (loaderstring: string): Loader {
  let [loader, query] = loaderstring.split('?')
  
  let L = loaders[loader]
  if (L != null) {
    return L(querystring.parse(query))
  }

  return null
}
