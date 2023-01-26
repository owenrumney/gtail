package logs

import (
	"github.com/owenrumney/gtail/pkg/output"
)

type LogAccess struct {
	projectID            string
	interestedSeverities []string
	outputWriter         output.Output
}

func New(projectID string, interestedSeverities []string, outputWriter output.Output) *LogAccess {
	return &LogAccess{
		projectID:            projectID,
		interestedSeverities: interestedSeverities,
		outputWriter:         outputWriter,
	}
}

func (la *LogAccess) Write(value interface{}) error {
	return la.outputWriter.Write(value)
}
