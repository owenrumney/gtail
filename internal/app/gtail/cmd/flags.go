package cmd

import (
	"os"
)

var (
	projectID        string = os.Getenv("GCP_PROJECT_ID")
	region           string = os.Getenv("GCP_REGION")
	logID            string
	severities       []string
	outputFormat     string
	debug            bool
	hoursAgo         int = 24
	buildTriggerName string
	functionName     string
	serviceName      string
	clusterName      string
	tailDuration     string = "10m"
	tailTopic        string
)

func init() {
	// need to set a default region
	if region == "" {
		region = "us-central1"
	}
}
