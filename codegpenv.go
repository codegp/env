package env

import "os"

var (
	isLocal         bool
	gcloudProjectID string
	dataStoreHost   string
)

// Lang is a type to represent a language supported by codegp
// the value must be the file extenstion for the language
type Lang string

const (
	// LocalStorePath simulates cloud storage by saving files to a local volume when running locally
	LocalStorePath string = "/localstore"
	// SourcePath is where the source downloader dowloads the projects source files
	SourcePath string = "/source"

	/*
		supported languages
	*/

	// Go is a Lang constant for the go language
	Go Lang = "go"
	// Python is a Lang constant for the python language
	Python Lang = "py"
	// Java is a Lang constant for the java language
	Java Lang = "java"
)

func init() {
	isLocal = os.Getenv("IS_LOCAL") == "true"
	gcloudProjectID = os.Getenv("GCLOUD_PROJECT_ID")
	dataStoreHost = os.Getenv("DATASTORE_EMULATOR_HOST")
}

// IsLocal returns true if running locally
func IsLocal() bool {
	return isLocal
}

// GCloudProjectID returns the google cloud project identifer
func GCloudProjectID() string {
	return gcloudProjectID
}

// DataStoreHost returns the datastore host
func DataStoreHost() string {
	return dataStoreHost
}
