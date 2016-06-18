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

interface UploadSummaryForm {
  SummaryUpload: SummaryUpload,
  Attachments: { [localPath: string]: fs.ReadStream }
}

const uploadPath = '/summary/upload'
const protocol = 'http://'

function UploadSummary(host: string, auth: string, summary: SummaryUpload, handler: (err, response: UploadSummaryResponse) => any) {
  let postData = summary

  let formData = {
    SummaryUpload: summary,
    Attachments: {},
  }

  summary.Artifacts.forEach((artUpload) => {
    formData.Attachments[artUpload.Path] = fs.createReadStream(artUpload.Path)
  })

  let requestData: request.Options = {
    url: protocol + host + uploadPath,
    headers: {
      'Authorization-Key': auth,
    },
    formData: formData
  }

  request.post(requestData, (err, response, body) => {
    handler(err, body)
  })
}

export { UploadSummary }
