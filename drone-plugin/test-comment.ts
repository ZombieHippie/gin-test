
import { PostComment } from './lib-ts/github/post-comment.function'

//PostComment('config/protractor.conf.js', 11, ":cloud:", 0)
//PostComment('config/protractor.conf.js', 12, ":horse:", 1)

import { LoadSettings } from './load-settings.function'

let settings = LoadSettings()

Post(`
http://cov.dryclean.io/data
`)


function Post(message: string) {
  PostComment(settings.github_user_handle, settings.github_user_password,
    message, "ZombieHippie/fantastic-repository", "e3598d42bfbd4b99bee3b550f983784ca187abcd",
    null, null, (err, resp) => {
      console.log("POST v ERR >", err)
      console.log(resp)
    })
}

