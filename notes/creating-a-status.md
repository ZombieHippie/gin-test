## Creating a status
https://developer.github.com/v3/repos/statuses/#create-a-status

Users with push access can create commit statuses for a given ref:

POST /repos/:owner/:repo/statuses/:sha
Note: there is a limit of 1000 statuses per sha and context within a Repository. Attempts to create more than 1000 statuses will result in a validation error.

### Parameters
```
Name        Type	    Description
state	      string	  Required. The state of the status. Can be one of pending, success, error, or failure.
target_url	string	  The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the 'source' of the Status.
  For example, if your Continuous Integration system is posting build status, you would want to provide the deep link for the build output for this specific SHA:
  http://ci.example.com/user/repo/build/sha.
description	string  	A short description of the status.
context	    string  	A string label to differentiate this status from the status of other systems. Default: "default"
```

### Example
```json
{
  "state": "success",
  "target_url": "https://example.com/build/status",
  "description": "The build succeeded!",
  "context": "continuous-integration/jenkins"
}
```
#### Response

Status: 201 Created
Location: https://api.github.com/repos/octocat/Hello-World/statuses/1
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4999
```json
{
  "created_at": "2012-07-20T01:19:13Z",
  "updated_at": "2012-07-20T01:19:13Z",
  "state": "success",
  "target_url": "https://ci.example.com/1000/output",
  "description": "Build has completed successfully",
  "id": 1,
  "url": "https://api.github.com/repos/octocat/Hello-World/statuses/1",
  "context": "continuous-integration/jenkins",
  "creator": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  }
}
```
