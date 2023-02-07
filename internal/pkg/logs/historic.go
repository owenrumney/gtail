package logs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/logging/logadmin"
	"github.com/googleapis/gax-go/v2/apierror"
	"github.com/owenrumney/gtail/internal/pkg/auth"
	"github.com/owenrumney/gtail/internal/pkg/logfilter"
	"google.golang.org/api/iterator"
)

func (la *LogAccess) GetHistoricalLogEntries(logFilter *logfilter.LogFilter) error {
	filterString, err := logFilter.GetFilterString()
	if err != nil {
		return err
	}

	adminClient, err := logadmin.NewClient(context.Background(), la.projectID)
	if err != nil {
		auth.CheckErrorForAuth(err)
		return err
	}
	defer func() { _ = adminClient.Close() }()
	iter := adminClient.Entries(context.Background(),
		logadmin.Filter(filterString),
	)

	for {
		entry, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			var apiErr *apierror.APIError
			if errors.As(err, &apiErr) {
				if apiErr.GRPCStatus().Code().String() == "ResourceExhausted" {
					fmt.Print("\nYou've hit the rate limit for requests... try limiting the search a bit or raise an issue at https://github.com/owenrumney/gtail\n")
					os.Exit(1)
				}
			}
		}
		if len(la.interestedSeverities) > 0 && entry != nil {
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
