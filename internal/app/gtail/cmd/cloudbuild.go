package cmd

import (
	"fmt"
	"strings"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/owenrumney/gtail/internal/pkg/logfilter"
	"github.com/owenrumney/gtail/internal/pkg/logs"
	"github.com/owenrumney/gtail/pkg/logger"
	"github.com/owenrumney/gtail/pkg/output"
	"github.com/spf13/cobra"
)

var cloudBuild = &cobra.Command{
	Use:     "cloud-build",
	Aliases: []string{"cb"},
	Short:   "Tail logs for a CloudBuild Job",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudBuildLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudBuildLogFilterType).
			WithBuildTriggerName(buildTriggerName).
			WithID(logID)
		return la.StreamLogEntries(lf)
	},
}

var historicCloudBuildCmd = &cobra.Command{
	Use:   "historic",
	Short: "Get the cloud build logs for a trigger that has already completed",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudBuildLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudBuildLogFilterType).
			WithBuildTriggerName(buildTriggerName).
			WithID(logID).
			WithHoursAgo(hoursAgo)
		return la.GetHistoricalLogEntries(lf)
	},
}

func getCloudBuildCommand() *cobra.Command {
	// add the child commands
	cloudBuild.AddCommand(historicCloudBuildCmd)

	// add the flags
	cloudBuild.PersistentFlags().StringVar(&buildTriggerName, "trigger-name", buildTriggerName, "The name of the cloud build trigger to use")
	cloudBuild.PersistentFlags().StringVar(&logID, "build-id", logID, "The cloud build ID")
	cloudBuild.PersistentFlags().StringSliceVar(&severities, "severity", severities, "The severity of logs to include")
	cloudBuild.PersistentFlags().StringVarP(&outputFormat, "output", "o", outputFormat, "The output format either json or a template string")
	historicCloudBuildCmd.Flags().IntVar(&hoursAgo, "hours-ago", hoursAgo, "Roughly how many hours ago the build happened. Searches a window of time from then till now")

	return cloudBuild
}

func defaultCloudBuildLogWriter(value interface{}) error {
	switch t := value.(type) {

	case *loggingpb.LogEntry:
		var content string
		switch payload := t.Payload.(type) {

		case *loggingpb.LogEntry_TextPayload:
			content = payload.TextPayload
		case *loggingpb.LogEntry_ProtoPayload:
			content = fmt.Sprintf("%v", payload.ProtoPayload)
		case *loggingpb.LogEntry_JsonPayload:
			content = fmt.Sprintf("%v", payload.JsonPayload)
		}
		if step, ok := t.Labels["build_step"]; ok {
			content = strings.TrimPrefix(strings.TrimPrefix(content, step), ": ")
		}
		fmt.Printf("%v\n", content)
	case *logging.Entry:
		content := t.Payload.(string)
		if step, ok := t.Labels["build_step"]; ok {
			content = strings.TrimPrefix(strings.TrimPrefix(content, step), ": ")
		}
		fmt.Printf("%v\n", content)
	default:
		logger.Debug("Got a default type: %v", t)
		fmt.Printf("%v\n", value)
	}
	return nil
}
