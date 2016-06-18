/// <reference path="../../typings/index.d.ts" />
import http = require("http")
import request = require("request")
import fs = require('fs')
import { Summary } from "../summary/summary.model"
import { Artifact } from "../artifact/artifact.model"
import { SummaryUpload } from "../summary/summary-upload.model"

export interface UploadSummaryResponse {
	Message: string
	Summary: Summary
  Artifacts: Artifact[]
  Count: number
}

interface Attachments {
  [localPath: string]: fs.ReadStream | string
}

interface UploadSummaryForm {
  SummaryUpload: SummaryUpload,
  Attachments: Attachments
}

const uploadPath = '/summary/upload'
const protocol = 'http://'

function UploadSummary(host: string, auth: string, summary: SummaryUpload, handler: (err, response: UploadSummaryResponse) => any) {

  let uploadFiles: Attachments = {}

  summary.Artifacts.forEach((artUpload) => {
    uploadFiles[artUpload.Path] = fs.createReadStream(artUpload.Path)
  })

  uploadFiles['SummaryUpload'] = JSON.stringify(summary)

  let requestData: request.Options = {
    url: protocol + host + uploadPath,
    method: 'POST',
    headers: {
      'Authorization-Key': auth,
    },
    formData: uploadFiles,
  }

  request(requestData, (err, response, body) => {
    handler(err, body)
  })
}

export { UploadSummary }
