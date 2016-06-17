interface DroneRepo {
  clone_url: string // "git://github.com/drone/drone"
  owner: string // "drone"
  name: string // "drone"
  full_name: string // "drone/drone"
}

interface DroneSystem {
  link_url: string // "https://beta.drone.io"
}

interface DroneBuild {
  number: number // 22,
  status: string // "success",
  started_at: number // 1421029603,
  finished_at: number // 1421029813,
  message: string // "Update the Readme",
  author: string // "johnsmith",
  author_email: string // "john.smith@gmail.com"
  event: string // "push",
  branch: string // "master",
  commit: string // "436b7a6e2abaddfd35740527353e78a227ddcb2c",
  ref: string // "refs/heads/master"
}

interface DroneWorkspace {
  root: string // "/drone/src",
  path: string // "/drone/src/github.com/drone/drone"
}

interface DroneVargs {
  
}

interface DroneParams {
  vargs: DroneVargs
  workspace: DroneWorkspace
  build: DroneBuild
  repo: DroneRepo
}
