package cmd

import (
	"os"
)

var (
	projectID        string = os.Getenv("GCP_PROJECT_ID")
	region           string = "us-central1"
	logID            string
	severities       []string
	outputFormat     string
	debug            bool
	hoursAgo         int = 24
	buildTriggerName string
	tailDuration     string = "10m"
	tailTopic        string
	serviceName      string
)
