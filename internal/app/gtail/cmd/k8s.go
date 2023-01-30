package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/owenrumney/gtail/internal/pkg/logfilter"
	"github.com/owenrumney/gtail/internal/pkg/logs"
	"github.com/owenrumney/gtail/pkg/logger"
	"github.com/owenrumney/gtail/pkg/output"
	"github.com/spf13/cobra"
	appEngine "google.golang.org/genproto/googleapis/appengine/logging/v1"
	"google.golang.org/genproto/googleapis/cloud/audit"
	"google.golang.org/protobuf/proto"
)

var k8sCommand = &cobra.Command{
	Use:   "k8s",
	Short: "Tail logs for GKE Cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultK8sLogWriter)
		la := logs.New(projectID, severities, outputWriter)

		lf := logfilter.New(projectID, logfilter.K8sClusterLogFilterType).
			WithServiceName(serviceName).
			WithRegion(region)

		return la.StreamLogEntries(lf)
	},
}

func getK8sCommand() *cobra.Command {
	// add the flags
	k8sCommand.PersistentFlags().StringVar(&clusterName, "cluster", clusterName, "GKE Cluster you want to tail logs for")
	k8sCommand.PersistentFlags().StringSliceVar(&severities, "severity", severities, "The severity of logs to include")
	k8sCommand.PersistentFlags().StringVarP(&outputFormat, "output", "o", outputFormat, "The output format either json or a template string")

	return k8sCommand
}

func defaultK8sLogWriter(value interface{}) error {
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
			switch payload.ProtoPayload.TypeUrl {
			case "type.googleapis.com/google.cloud.audit.AuditLog":
				var auditLog audit.AuditLog
				err := proto.Unmarshal(payload.ProtoPayload.Value, &auditLog)
				if err != nil {
					logger.Error("error unmarshalling audit log: %v", err)
					content = fmt.Sprintf("%v", payload.ProtoPayload)
				} else {
					contentBytes, err := json.MarshalIndent(auditLog, "", "  ")
					if err != nil {
						logger.Error("error marshalling audit log: %v", err)
						content = fmt.Sprintf("%v", auditLog)
					} else {
						content = string(contentBytes)
					}
				}
			case "type.googleapis.com/google.appengine.logging.v1.RequestLog":
				var requestLog appEngine.RequestLog
				err := proto.Unmarshal(payload.ProtoPayload.Value, &requestLog)
				if err != nil {
					logger.Error("error unmarshalling audit log: %v", err)
					content = fmt.Sprintf("%v", payload.ProtoPayload)
				} else {
					content = fmt.Sprintf("%v", requestLog)
				}
			default:
				content = fmt.Sprintf("%v", payload.ProtoPayload)
			}

		case *loggingpb.LogEntry_JsonPayload:
			c, err := json.MarshalIndent(payload.JsonPayload, "", "  ")
			if err != nil {
				logger.Error("error marshalling json payload: %v", err)
				content = fmt.Sprintf("%v", payload.JsonPayload)
			} else {
				content = string(c)
			}
		}

		fmt.Printf("%s: %v\n", t.Timestamp.AsTime().Format(time.RFC3339), content)
	default:
		fmt.Printf("%v\n", value)
	}
	return nil
}
