# gitlab
GitLab API client for golang

# Tests

To run tests, you will need a working GitLab server.

The test code needs several info of that GitLab server, passing via environment variables:

```sh
# personal access token
export PERSONAL_TOKEN="supersecret"
# server address
export SERVER_URL="http://127.0.0.1:3333"
# api path, should be /api/v3
export API_PATH="/api/v3"
# repository name, the repository must be a mirror of
#   https://gitlab.com/Ronmi/test-project.git
# The easiest way to obtain one is importing it into
# your GitLab server.
export REPO_PATH="root/test"
# repository id, you can find it at "Trggers" page in your project settings
export REPO_ID=2

go test ./...
```
