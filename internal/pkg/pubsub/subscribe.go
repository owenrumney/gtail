package pubsub

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/owenrumney/gtail/internal/pkg/auth"
	"github.com/owenrumney/gtail/pkg/logger"
)

func (t *PubSub) createSubscription(subscriptionName, topicName, pushEndpoint string) (*pubsub.Subscription, error) {
	if topicName == "" {
		return nil, fmt.Errorf("topic name is required")
	}
	if subscriptionName == "" {
		return nil, fmt.Errorf("subscription name is required")
	}

	logger.Debug("Creating subscription: %s", subscriptionName)
	subscription := t.client.Subscription(subscriptionName)
	exists, err := subscription.Exists(t.ctx)
	if err != nil {
		auth.CheckErrorForAuth(err)
		logger.Error("Error checking if subscription exists: %v", err)
	}
	if exists {
		logger.Warn("Topic already exists: %s", topicName)
		return subscription, nil
	}

	subscriptionConfig := pubsub.SubscriptionConfig{
		Topic: t.client.Topic(topicName),
	}

	if pushEndpoint != "" {
		subscriptionConfig.PushConfig = pubsub.PushConfig{
			Endpoint: pushEndpoint,
		}
	}

	ctx, cancel := context.WithTimeout(t.ctx, 10*time.Second)
	defer cancel()
	subscription, err = t.client.CreateSubscription(ctx, subscriptionName, subscriptionConfig)
	if err != nil {
		return nil, err
	}
	logger.Info("Subscription created: %s", subscriptionName)
	return subscription, nil
}

func (t *PubSub) deleteSubscription(subscriptionName string) error {
	logger.Debug("Deleting subscription: %s", subscriptionName)
	subscription := t.client.Subscription(subscriptionName)
	exists, err := subscription.Exists(t.ctx)
	if err != nil {
		auth.CheckErrorForAuth(err)
		logger.Error("Error checking if subscription exists: %v", err)
	}
	if !exists {
		logger.Warn("Subscription does not exist: %s", subscriptionName)
		return nil
	}

	if err := subscription.Delete(t.ctx); err != nil {
		return err
	}
	logger.Info("Subscription deleted: %s", subscriptionName)
	return nil
}

func (ps *PubSub) StreamSubscription(topicName, tailDuration string) error {

	subscriptionName := fmt.Sprintf("gtail-%s-%d", topicName, time.Now().Unix())
	logger.Info("Creating tail subscription: %s", subscriptionName)
	subscription, err := ps.createSubscription(subscriptionName, topicName, "")
	if err != nil {
		return err
	}
	defer subscription.Delete(ps.ctx)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("Deleting subscription before exiting: %s", subscriptionName)
		subscription.Delete(ps.ctx)
		os.Exit(0)
	}()

	ctx, cancel := context.WithTimeout(ps.ctx, 10*time.Second)
	defer cancel()
	_, err = subscription.Exists(ctx)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(tailDuration)
	if err != nil {
		logger.Error("Error parsing duration: %v", err)
		duration = 10 * time.Minute
	}

	logger.Info("Tailing topic: %s for %s", topicName, tailDuration)
	ctx, cancel = context.WithTimeout(ps.ctx, duration)
	defer cancel()

	return subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		// lock the output writer so we don't get interleaved output
		ps.Lock()
		defer ps.Unlock()
		ps.outputWriter.Write(msg)
	},
	)
}
