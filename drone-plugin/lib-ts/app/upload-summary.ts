/// <reference path="../../typings/index.d.ts" />
import http = require("http")
import request = require("request")
import fs = require('fs')
import path = require('path')
import { Summary } from "../summary/summary.model"
import { Artifact } from "../artifact/artifact.model"
import { SummaryUpload } from "../upload/summary-upload.model"

interface ZipFolder {
  (srcFolder: string, zipFilePath: string, callback: (err) => any): any
}
const zipFolder = require('zip-folder') as ZipFolder

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

  let zipsToUpload = 0

  summary.Artifacts.forEach((artUpload) => {
    try {
      let stats = fs.statSync(artUpload.Path)
      // if directory upload zip
      if (stats.isDirectory()) {
        console.log("Found directory to upload, uploading as zip")
        let zippedfilename = artUpload.Path.replace(/[^\w\-]+/g, '-')
        zippedfilename = ('artifact-' + zippedfilename).replace(/\-+/g, '-')
        addZipToQueue(artUpload.Path, zippedfilename + '.zip', artUpload.FormKey)
        artUpload.Archived = true
      } else {
        uploadFiles[artUpload.FormKey] = fs.createReadStream(artUpload.Path)
      }
    } catch (err) {
      console.error(`Error creating file stream for ${artUpload.Path}!`)
    }
  })

  uploadFiles['SummaryUpload'] = JSON.stringify(summary)


  // give a quick check just in case there were no directories to upload
  checkFinishedUploading()

  function addZipToQueue (src, dest, formKey: string) {
    zipsToUpload += 1
    zipFolder(src, dest, function (err) {
      if (err) {
        console.error(`Error writing zip file!`)
      } else {
        console.log("Success writing zip file")
        // Create the read stream for the form
        uploadFiles[formKey] = fs.createReadStream(dest)
      }
      zipsToUpload--
      checkFinishedUploading()
    })
  }

  // if no zip files to wait on, finish uploading
  function checkFinishedUploading (force = false) {
    if (zipsToUpload === 0 || force) {
      let requestData: request.Options = {
        url: protocol + host + uploadPath,
        method: 'POST',
        headers: {
          'Authorization-Key': auth,
        },
        formData: uploadFiles,
      }

      request(requestData, (err, response, body) => {
        try {
          body = JSON.parse(body)
        } catch (err) {}
        handler(err, body)
      })
    }
  }
}

export { UploadSummary }
