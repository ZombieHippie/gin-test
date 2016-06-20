/// <reference path="./typings/index.d.ts" />

interface settings {
  provider:   string // "github",

  github_client_id:     string // "APPLICATION_CLIENT_ID",
  github_client_secret: string // "APPLICATION_CLIENT_SECRET",
  github_user_handle:   string // "Username",
  github_user_password: string // "Password or Token"
}

export function LoadSettings (): settings {
  return require('./settings.json') as settings
}
