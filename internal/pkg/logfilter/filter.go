package logfilter

import (
	"fmt"
	"strings"
	"time"

	"github.com/owenrumney/gtail/pkg/logger"
)

type LogFilterType string

const (
	CloudBuildLogFilterType    LogFilterType = "build"
	CloudRunLogFilterType      LogFilterType = "run"
	CloudFunctionLogFilterType LogFilterType = "function"
	K8sClusterLogFilterType    LogFilterType = "k8s_cluster"
)

type LogFilter struct {
	filterType       LogFilterType
	projectID        string
	historic         bool
	hoursAgo         int
	logID            string
	buildTriggerName string
	functionName     string
	serviceName      string
	clusterName      string
	region           string
}

func New(projectID string, filterType LogFilterType) *LogFilter {
	return &LogFilter{
		filterType:       filterType,
		projectID:        projectID,
		historic:         false,
		hoursAgo:         24,
		logID:            "",
		buildTriggerName: "",
		functionName:     "",
		serviceName:      "",
		clusterName:      "",
		region:           "us-central1",
	}
}

func (lf *LogFilter) WithHistoric(historic bool) *LogFilter {
	lf.historic = historic
	return lf
}

func (lf *LogFilter) WithHoursAgo(hoursAgo int) *LogFilter {
	lf.hoursAgo = hoursAgo
	return lf
}

func (lf *LogFilter) WithID(id string) *LogFilter {
	lf.logID = id
	return lf
}

func (lf *LogFilter) WithBuildTriggerName(buildTriggerName string) *LogFilter {
	lf.buildTriggerName = buildTriggerName
	return lf
}

func (lf *LogFilter) WithFunctionName(functionName string) *LogFilter {
	lf.functionName = functionName
	return lf
}

func (lf *LogFilter) WithServiceName(serviceName string) *LogFilter {
	lf.serviceName = serviceName
	return lf
}

func (lf *LogFilter) WithClusterName(clusterName string) *LogFilter {
	lf.clusterName = clusterName
	return lf
}

func (lf *LogFilter) WithRegion(region string) *LogFilter {
	lf.region = region
	return lf
}

func (lf *LogFilter) GetFilterString() string {
	filters := []string{}

	switch lf.filterType {
	case CloudBuildLogFilterType:
		filters = append(filters, fmt.Sprintf(`log_name="projects/%s/logs/cloudbuild"`, lf.projectID))
		filters = append(filters, `resource.type="build"`)
		if lf.logID != "" {
			filters = append(filters, fmt.Sprintf(`resource.labels.build_id="%s"`, lf.logID))
		}
	case CloudRunLogFilterType:
		filters = append(filters, fmt.Sprintf(`log_name: "projects/%s/logs/run.googleapis.com"`, lf.projectID))
		filters = append(filters, `resource.type="cloud_run_revision"`)
		if lf.logID != "" {
			filters = append(filters, fmt.Sprintf(`resource.labels.revision_name="%s"`, lf.logID))
		}
	case CloudFunctionLogFilterType:
		filters = append(filters, `resource.type="cloud_function"`)
		if lf.logID != "" {
			filters = append(filters, fmt.Sprintf(`resource.labels.execution_id="%s"`, lf.logID))
		}
	case K8sClusterLogFilterType:
		filters = append(filters, `resource.type="k8s_cluster"`, fmt.Sprintf(`resource.labels.project_id="%s"`, lf.projectID))
	}

	if lf.region != "" {
		filters = append(filters, fmt.Sprintf(`resource.labels.location="%s"`, lf.region))
	}

	if lf.historic {
		start := time.Now().Add(time.Duration(-lf.hoursAgo) * time.Hour)
		end := start.Add(time.Duration(lf.hoursAgo) * time.Hour)

		filters = append(filters, fmt.Sprintf(`timestamp>="%s"`, start.Format(time.RFC3339)))
		filters = append(filters, fmt.Sprintf(`timestamp<="%s"`, end.Format(time.RFC3339)))
	}

	if lf.serviceName != "" {
		filters = append(filters, fmt.Sprintf(`resource.labels.service_name="%s"`, lf.serviceName))
	}

	if lf.clusterName != "" {
		filters = append(filters, fmt.Sprintf(`resource.labels.cluster_name="%s"`, lf.clusterName))
	}

	if lf.functionName != "" {
		filters = append(filters, fmt.Sprintf(`resource.labels.function_name="%s"`, lf.functionName))
	}

	if lf.buildTriggerName != "" {
		triggerID, err := resolveBuildTriggerID(lf.projectID, lf.buildTriggerName)
		if err != nil {
			logger.Error("could not resolve the trigger named %s: %v", lf.buildTriggerName, err)
		}
		filters = append(filters, fmt.Sprintf(`resource.labels.build_trigger_id="%s"`, triggerID))
	}

	filterString := strings.Join(filters, " ")
	logger.Debug("Using filter string [%s]", filterString)
	return filterString
}
