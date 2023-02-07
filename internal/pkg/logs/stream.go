package logs

import (
	"context"
	"fmt"
	"io"
	"strings"

	logging "cloud.google.com/go/logging/apiv2"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/owenrumney/gtail/internal/pkg/auth"
	"github.com/owenrumney/gtail/internal/pkg/logfilter"
	"github.com/owenrumney/gtail/pkg/logger"
)

func (la *LogAccess) StreamLogEntries(logFilter *logfilter.LogFilter) error {
	ctx := context.Background()
	client, err := logging.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("NewClient error: %v", err)
	}
	defer func() { _ = client.Close() }()

	stream, err := client.TailLogEntries(ctx)
	if err != nil {
		auth.CheckErrorForAuth(err)
		return fmt.Errorf("TailLogEntries error: %v", err)
	}
	defer func() { _ = stream.CloseSend() }()

	filter, err := logFilter.GetFilterString()
	if err != nil {
		return err
	}

	req := &loggingpb.TailLogEntriesRequest{
		ResourceNames: []string{
			fmt.Sprintf("projects/%s", la.projectID),
		},
		Filter: filter,
	}
	if err := stream.Send(req); err != nil {
		return fmt.Errorf("stream.Send error: %v", err)
	}

	// read and print two or more streamed log entries
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			logger.Debug("stream.Recv EOF")
			break
		}
		if err != nil {
			return fmt.Errorf("stream.Recv error: %v", err)
		}
		if resp.Entries != nil {
			for _, entry := range resp.Entries {
				if len(la.interestedSeverities) > 0 {
					for _, s := range la.interestedSeverities {
						if strings.EqualFold(entry.Severity.String(), s) {
							if err := la.Write(entry); err != nil {
								logger.Error("error writing the log entry %v", err)
							}
							break
						}
					}
				} else {
					if err := la.Write(entry); err != nil {
						logger.Error("error writing the log entry %v", err)
					}
				}
			}
		}
	}
	return nil
}
