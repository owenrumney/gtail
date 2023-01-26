package pubsub

import (
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/owenrumney/gtail/internal/pkg/auth"
	"github.com/owenrumney/gtail/pkg/output"
)

type PubSub struct {
	sync.Mutex
	ctx          context.Context
	client       *pubsub.Client
	outputWriter output.Output
}

func New(projectID string, outputWriter output.Output) (*PubSub, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		auth.CheckErrorForAuth(err)
		return nil, err
	}

	return &PubSub{
		ctx:          ctx,
		client:       client,
		outputWriter: outputWriter,
	}, nil
}
