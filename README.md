# Jenkins Test Repository

This is a test repository for validating Frogbot's CI auto-detection feature with Jenkins.

## What's in here?

- **Go application**: Simple REST API using Gin framework
- **Dependencies**: gin-gonic/gin, sirupsen/logrus, gopkg.in/yaml.v3
- **Jenkinsfile**: Jenkins pipeline that tests Frogbot auto-detection

## Frogbot Auto-Detection Test

The Jenkinsfile will test that Frogbot can automatically detect:
- `JF_GIT_OWNER` - from GIT_URL
- `JF_GIT_REPO` - from GIT_URL
- `JF_GIT_BASE_BRANCH` - from GIT_LOCAL_BRANCH

## Running locally

```bash
go mod download
go run main.go
```

Then visit http://localhost:8080

## Jenkins Setup

1. Create a Multibranch Pipeline in Jenkins
2. Point it to this repository
3. Configure credentials:
   - `jfrog-url` - Your JFrog platform URL
   - `jfrog-token` - JFrog access token
   - `git-token` - GitHub personal access token
4. Run the pipeline!
