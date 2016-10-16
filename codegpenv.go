package codegpenv

import "os"

var (
	isLocal         bool
	gcloudProjectID string
)

const (
	localStorePath string = "/localstore"
	sourcePath     string = "/source"
)

func init() {
	isLocal = os.Getenv("IS_LOCAL") == "true"
	gcloudProjectID = os.Getenv("GCLOUD_PROJECT_ID")
}

func IsLocal() bool {
	return isLocal
}

func GCloudProjectID() string {
	return gcloudProjectID
}

func LocalStorePath() string {
	return localStorePath
}

func SourcePath() string {
	return sourcePath
}
