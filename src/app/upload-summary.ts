/// <reference path="../../typings/index.d.ts" />
import http = require("http")
import querystring = require("querystring")
import { Summary } from "../summary/summary.model"

function UploadSummary(host: string, auth: string, summary: Summary, handler: (err, summary: Summary) => any) {
  let postData = querystring.stringify(summary)
  let [hostname, port] = host.split(':')
  let requestData = {
    host: hostname,
    port: parseInt(port),
    path: '/summary/webhook',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Content-Length': Buffer.byteLength(postData),
      'Authorization-Key': auth,
    },
  }

  let postReq = http.request(requestData, (response) => {
    if (response.statusCode < 200 || response.statusCode > 299) {
      handler(Error("Error returned from server"), null)
      return
    } 
    response.setEncoding('utf8')
    let responseBody = ""
    response.on('data', (chunk) => {
        responseBody += chunk.toString()
    })
    response.on('end', () => {
      handler(null, JSON.parse(responseBody))
    })
  })

  postReq.write(postData)
  postReq.end()
}

export { UploadSummary }
