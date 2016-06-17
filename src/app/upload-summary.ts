/// <reference path="../../typings/index.d.ts" />
import http = require("http")
import request = require("request")
import { Summary } from "../summary/summary.model"

export interface UploadSummaryResponse {
	Message: string
	Summary: Summary
}

const uploadPath = '/summary/webhook'
const protocol = 'http://'

function UploadSummary(host: string, auth: string, summary: Summary, handler: (err, response: UploadSummaryResponse) => any) {
  let postData = summary
  let requestData: request.Options = {
    url: protocol + host + uploadPath,
    method: 'POST',
    headers: {
      'Authorization-Key': auth,
    },
    json: true,
    body: postData,
  }

  request(requestData, (err, response, body) => {
    console.log(err, body)

    handler(err, body)
    
  })
}

export { UploadSummary }
