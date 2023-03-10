package cmd

import (
	"fmt"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/owenrumney/gtail/internal/pkg/logfilter"
	"github.com/owenrumney/gtail/internal/pkg/logs"
	"github.com/owenrumney/gtail/pkg/output"
	"github.com/spf13/cobra"
)

var cloudRunCmd = &cobra.Command{
	Use:     "cloud-run",
	Aliases: []string{"cr"},
	Short:   "Tail logs for a Cloud Run revision",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudRunLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudRunLogFilterType).
			WithServiceName(serviceName).
			WithRegion(region)

		return la.StreamLogEntries(lf)
	},
}

var historicCloudRunCmd = &cobra.Command{
	Use:   "historic",
	Short: "Get the cloud run logs for a revision that has already exited",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudRunLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudRunLogFilterType).
			WithServiceName(serviceName).
			WithID(logID)
		return la.GetHistoricalLogEntries(lf)
	},
}

func getCloudRunCommand() *cobra.Command {
	// add the child commands
	cloudRunCmd.AddCommand(historicCloudRunCmd)

	// add the flags
	cloudRunCmd.PersistentFlags().StringVar(&serviceName, "service", "", "Cloud Run service for the logs to get")
	cloudRunCmd.PersistentFlags().StringVar(&logID, "revision-id", logID, "The cloud run revision ID")
	cloudRunCmd.PersistentFlags().StringSliceVar(&severities, "severity", severities, "The severity of logs to include")
	cloudRunCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", outputFormat, "The output format either json or a template string")
	historicCloudRunCmd.Flags().IntVar(&hoursAgo, "hours-ago", hoursAgo, "Roughly how many hours ago the log happened. Searches a window of time from then till now")

	return cloudRunCmd
}

func defaultCloudRunLogWriter(value interface{}) error {
	switch t := value.(type) {
	case *logging.Entry:
		if content, ok := t.Payload.(string); ok {
			fmt.Printf("%v\n", content)
		}
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
		fmt.Printf("%v\n", content)
	default:
		fmt.Printf("%v\n", value)
	}
	return nil
}
