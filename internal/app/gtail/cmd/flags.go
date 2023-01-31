package cmd

import (
	"os"
)

var (
	debug            bool
	lastRun          bool
	hoursAgo         int    = 24
	projectID        string = os.Getenv("GCP_PROJECT_ID")
	region           string = os.Getenv("GCP_REGION")
	logID            string
	outputFormat     string
	buildTriggerName string
	functionName     string
	serviceName      string
	clusterName      string
	tailDuration     string = "10m"
	tailTopic        string
	severities       []string
)

func init() {
	// need to set a default region
	if region == "" {
		region = "us-central1"
	}
}
