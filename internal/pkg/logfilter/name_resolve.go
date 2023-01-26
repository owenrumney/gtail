package logfilter

import (
	"context"
	"fmt"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1/v2"
	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/owenrumney/gtail/pkg/logger"
)

func resolveBuildTriggerID(projectID, triggerName string) (string, error) {
	client, err := cloudbuild.NewClient(context.Background())
	if err != nil {
		return "", fmt.Errorf("NewClient error: %v", err)
	}
	defer func() { _ = client.Close() }()

	trigger, err := client.GetBuildTrigger(context.Background(), &cloudbuildpb.GetBuildTriggerRequest{
		ProjectId: projectID,
		TriggerId: triggerName,
	})
	if err != nil {
		return "", fmt.Errorf("GetBuildTrigger error: %v", err)
	}
	logger.Debug("Resolved trigger ID for %s: %s", trigger, trigger.Id)
	return trigger.Id, nil
}
