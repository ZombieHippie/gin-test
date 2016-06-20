/// <reference path="../../typings/index.d.ts" />
/**
 * From notes.

## Creating a Commit Comment
https://developer.github.com/v3/repos/comments/#create-a-commit-comment
POST /repos/:owner/:repo/commits/:sha/comments
### Input
```
Name      Type        Description
body      string      Required. The contents of the comment.
path      string      Relative path of the file to comment on.
position  integer	    Line index in the diff to comment on.
line      integer	    Deprecated. Use position parameter instead. Line number in the file to comment on.
```
### Example
{
  "body": "Great stuff",
  "path": "file1.txt",
  "position": 4,
  "line": null
}
 
 */


import request = require("request")
import fs = require('fs')
import { ENV } from '../../drone-parser'

interface CommentPayload {
  body:     string, // Required. The contents of the comment.
  path:     string, // Relative path of the file to comment on.
  position: number, // Line index in the diff to comment on.
}

const emoji = [
  ":hocho:"
]

function PostComment(botusername: string, botpassword: string, message: string,
                     repo: string,        sha: string,         path?: string,
                     position?: number,   callback: (err, response) => any = null) {

  const host = `https://${botusername}:${botpassword}@api.github.com`
  const endpoint = `/repos/${repo}/commits/${sha}/comments`

  let options: request.OptionsWithUrl = {
    url: host + endpoint,
    headers: {
      "User-Agent": "Drycleaner robot",
    },
    method: 'POST',
    json: {
      body: message,
      path: path,
      position: position,
    },
  }

  request(options, (error, httpresponse, body) => {
    console.log(error, body)
    try {
      body = JSON.parse(body)
    } catch (err) {} // guess it wasn't json...
    callback && callback(error, body)
  })
}

export { PostComment }
