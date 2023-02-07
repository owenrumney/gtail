package main

import (
	"os"
	"runtime/debug"

	"github.com/owenrumney/gtail/internal/app/gtail/cmd"
	"github.com/owenrumney/gtail/pkg/logger"
)

func main() {
	rootCmd := cmd.GetRootCmd()

	defer func() {
		if r := recover(); r != nil {

			logger.Error("Recovered from a fatal error: %#v. %s", r, string(debug.Stack()))
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
