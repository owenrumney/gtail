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

var cloudFunctionCmd = &cobra.Command{
	Use:     "cloud-function",
	Aliases: []string{"cf"},
	Short:   "Tail logs for a Cloud Function revision",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudFunctionLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudRunLogFilterType).
			WithFunctionName(functionName).
			WithRegion(region)

		return la.StreamLogEntries(lf)
	},
}

var historicCloudFunctionCmd = &cobra.Command{
	Use:   "historic",
	Short: "Get the cloud function logs for a run that has already exited",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultCloudRunLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.CloudFunctionLogFilterType).
			WithFunctionName(functionName)

		return la.GetHistoricalLogEntries(lf)
	},
}

func getCloudFunctionCommand() *cobra.Command {
	// add the child commands
	cloudFunctionCmd.AddCommand(historicCloudFunctionCmd)

	// add the flags
	cloudFunctionCmd.PersistentFlags().StringVar(&functionName, "function", "", "Cloud Function name for the logs to get")
	cloudFunctionCmd.PersistentFlags().StringVar(&logID, "execution-id", logID, "The Cloud Function execution ID")
	cloudFunctionCmd.PersistentFlags().StringSliceVar(&severities, "severity", severities, "The severity of logs to include")
	cloudFunctionCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", outputFormat, "The output format either json or a template string")
	historicCloudFunctionCmd.Flags().IntVar(&hoursAgo, "hours-ago", hoursAgo, "Roughly how many hours ago the log happened. Searches a window of time from then till now")

	return cloudFunctionCmd
}

func defaultCloudFunctionLogWriter(value interface{}) error {
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
