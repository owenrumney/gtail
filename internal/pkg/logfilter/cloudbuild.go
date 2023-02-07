package logfilter

import (
	"context"
	"fmt"
	"time"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1/v2"
	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/owenrumney/gtail/internal/pkg/auth"
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
		auth.CheckErrorForAuth(err)
		return "", fmt.Errorf("GetBuildTrigger error: %v", err)
	}
	logger.Debug("Resolved trigger ID for %s: %s", trigger, trigger.Id)
	return trigger.Id, nil
}

func getLatestBuildID(projectID, triggerID string) (string, *time.Time, error) {
	client, err := cloudbuild.NewClient(context.Background())
	if err != nil {
		return "", nil, fmt.Errorf("NewClient error: %v", err)
	}
	defer func() { _ = client.Close() }()

	logger.Debug("Getting latest build for trigger %s", triggerID)
	builds := client.ListBuilds(context.Background(), &cloudbuildpb.ListBuildsRequest{
		ProjectId: projectID,
		Filter:    fmt.Sprintf("trigger_id=%s", triggerID),
		PageSize:  1,
	})

	build, err := builds.Next()
	if err != nil {
		auth.CheckErrorForAuth(err)
		return "", nil, fmt.Errorf("ListBuilds error: %v", err)
	}
	createTime := build.GetCreateTime().AsTime()

	return build.Id, &createTime, nil

}
