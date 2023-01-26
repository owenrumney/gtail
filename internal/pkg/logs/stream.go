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

	req := &loggingpb.TailLogEntriesRequest{
		ResourceNames: []string{
			fmt.Sprintf("projects/%s", la.projectID),
		},
		Filter: logFilter.GetFilterString(),
	}
	if err := stream.Send(req); err != nil {
		return fmt.Errorf("stream.Send error: %v", err)
	}

	// read and print two or more streamed log entries
	for counter := 0; counter < 2; {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("stream.Recv error: %v", err)
		}
		if resp.Entries != nil {
			counter += len(resp.Entries)
			for _, entry := range resp.Entries {
				if len(la.interestedSeverities) > 0 {
					for _, s := range la.interestedSeverities {
						if strings.EqualFold(entry.Severity.String(), s) {
							if err := la.Write(entry); err != nil {
								return err
							}
							break
						}
					}
				} else {
					if err := la.Write(entry); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
