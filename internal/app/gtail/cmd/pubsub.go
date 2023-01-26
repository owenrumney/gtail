package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	gpubsub "cloud.google.com/go/pubsub"
	"github.com/owenrumney/gtail/internal/pkg/pubsub"
	"github.com/owenrumney/gtail/pkg/logger"
	"github.com/owenrumney/gtail/pkg/output"
	"github.com/spf13/cobra"
)

var pubsubCmd = &cobra.Command{
	Use:     "pubsub",
	Aliases: []string{"ps"},
	Short:   "Tail a pubsub topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		outputWriter := output.New(outputFormat, severities, defaultPubsubLogWriter)
		ps, err := pubsub.New(projectID, outputWriter)
		if err != nil {
			return err
		}

		return ps.StreamSubscription(tailTopic, tailDuration)
	},
}

func getPubSubCommand() *cobra.Command {

	pubsubCmd.Flags().StringVarP(&tailTopic, "topic", "t", "", "The pubsub topic to tail")
	pubsubCmd.Flags().StringVar(&tailDuration, "tail-duration", tailDuration, "The duration to tail for")
	pubsubCmd.Flags().StringVarP(&outputFormat, "output", "o", outputFormat, "The output format json or a template string")

	return pubsubCmd
}

func defaultPubsubLogWriter(value interface{}) error {

	if value == nil {
		logger.Warn("No message received")
		return nil
	}

	if msg, ok := value.(*gpubsub.Message); ok {

		eventHeader := fmt.Sprintf("Published: %s ID: %s", msg.PublishTime, msg.ID)
		fmt.Printf("\n%s\n%s\n", eventHeader, strings.Repeat("-", len(eventHeader)))
		for attr, val := range msg.Attributes {
			fmt.Printf("%s:\t\t%v\n", attr, val)
		}

		var content map[string]interface{}
		err := json.Unmarshal(msg.Data, &content)
		if err != nil {
			fmt.Printf("%s\n", string(msg.Data))
		}
		body, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			logger.Error("Error marshalling message: %v", err)
			fmt.Printf("%s\n", string(msg.Data))
		}
		fmt.Printf("%v\n", string(body))
	} else {
		fmt.Printf("%v\n", value)
	}
	return nil
}
