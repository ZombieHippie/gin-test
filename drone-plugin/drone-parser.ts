/// <reference path="./typings/index.d.ts" />

interface VEnv {
  PLUGIN_HOST?:      string // 'cov.dryclean.io',
  PLUGIN_FILES?:     VFile[] // '[{"label":"jasmine-coverage.json","path":"./coverage/Phantom*/coverage-final.json","postprocessor":"cobertura"},{"label":"tslint.txt","path":"./tslint-checkstyle.txt","postprocessor":"checkstyle"}]',
  PLUGIN_CONTEXT?:   string // 'drone-dryclean',
  PLUGIN_DEBUG?:     string // 'true',
  PLUGIN_AUTH?:      string // 'super-secret-auth',
  PLUGIN_REPO?:      string // 'colelawr/drone-dryclean',
}

interface VFile {
  path:           string // "stats.json", "coverage/Phantom*/index.html"
  label:          string // "webpack", "coverage"
  postprocessor?: string // "cobertura", "codestyle", etc
}


interface DRONE_ENV {
  CI?:       string // 'drone',
  HOME?:     string // '/root'
  PATH?:     string // '/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin',
  HOSTNAME?: string // '79a50f0de7ea',

  DRONE_BUILD_EVENT?:     string // 'push',
  DRONE_REPO_TRUSTED?:    boolean // 'false',
  DRONE?:                 boolean // 'true',
  DRONE_COMMIT_REF?:      string // 'refs/heads/master',
  DRONE_YAML_SIGNED?:     string // 'false',
  DRONE_REPO_OWNER?:      string // 'ZombieHippie',

  DRONE_COMMIT_AUTHOR_EMAIL?: string // 'msgzht@gmail.com',
  DRONE_COMMIT_SHA?:          string // '0b83add03af7193650c3f82a4dca4f458bf92509',
  DRONE_COMMIT_MESSAGE?:      string // 'Trying more options',
  
  
  DRONE_REPO_BRANCH?:  string // 'master',
  DRONE_COMMIT_LINK?:  string // 'https://github.com/ZombieHippie/fantastic-repository/commit/0b83add03af7193650c3f82a4dca4f458bf92509',
  DRONE_REPO_AVATAR?:  string // 'https://avatars.githubusercontent.com/u/2925395?v=3',
  DRONE_BRANCH?:       string // 'master',
  DRONE_REMOTE_URL?:   string // 'https://github.com/ZombieHippie/fantastic-repository.git',
  DRONE_BUILD_NUMBER?: number // '14',
  DRONE_ARCH?:         string // 'linux/amd64',
  DRONE_VERSION?:      string // '0.5.0+',
  DRONE_REPO_LINK?:    string // 'https://github.com/ZombieHippie/fantastic-repository',
  DRONE_REPO?:         string // 'ZombieHippie/fantastic-repository',

  
  DRONE_PREV_BUILD_NUMBER?: number  // '13',
  DRONE_PREV_COMMIT_SHA?:   string  // 'ef17bd452cf451ea1a3ae8a8672e3a8d31f501f6',
  DRONE_REPO_PRIVATE?:      boolean // 'false',
  DRONE_REPO_SCM?:          string  // 'git',
  DRONE_COMMIT_BRANCH?:     string  // 'master',
  DRONE_YAML_VERIFIED?:     boolean // 'false',
  DRONE_REPO_NAME?:         string  // 'fantastic-repository',

  DRONE_COMMIT_AUTHOR_AVATAR?: string // 'https://avatars.githubusercontent.com/u/2925395?v=3',
  DRONE_COMMIT_AUTHOR?:        string // 'ZombieHippie',

  DRONE_BUILD_STATUS?:      string // 'success',
  DRONE_PREV_BUILD_STATUS?: string // 'failure',
  DRONE_BUILD_LINK?:        string // 'http://drone.dryclean.io/ZombieHippie/fantastic-repository/14',
  DRONE_BUILD_CREATED?:     number // '1466304194',
  DRONE_BUILD_FINISHED?:    number // '0',
  DRONE_BUILD_STARTED?:     number // '0',
  DRONE_COMMIT?:            string // '0b83add03af7193650c3f82a4dca4f458bf92509',
}

// Union between drone and plugin specific vars
type PLUGIN_ENV = DRONE_ENV & VEnv

let drone_env: PLUGIN_ENV = {}

for (let key in process.env) {
  let val = process.env[key]
  try {
    val = JSON.parse(val)
  } catch (err) {} // remains a string
  
  drone_env[key] = val
}

export { drone_env as ENV, PLUGIN_ENV, VFile }
