package cmd

import (
	"fmt"
	"os"

	"github.com/owenrumney/gtail/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gtail",
	Short: "Tail GCP output for various services",
	Long:  "Tail the output from a number of different GCP services - PubSub, CloudBuild, CloudRun, etc.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" {
			return nil
		}

		if projectID == "" && os.Getenv("GCP_PROJECT_ID") == "" {
			fmt.Print("\nYou must specify a project ID using the -p flag or GCP_PROJECT_ID envvar.\n\n")
			_ = cmd.Usage()
			os.Exit(0)
		}

		logger.Configure(os.Stdout, debug)
		return nil
	},
}

func GetRootCmd() *cobra.Command {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceUsage = true

	rootCmd.AddCommand(getCloudBuildCommand(), getCloudRunCommand(), getPubSubCommand())
	rootCmd.PersistentFlags().StringVarP(&projectID, "project", "p", projectID, "The GCP project ID")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", region, "The GCP region")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	return rootCmd
}
