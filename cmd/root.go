package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

// rootCmd is the root for all commands
var rootCmd = &cobra.Command{
	Use:   "webapp-sidecar",
	Short: "WebApp Sidecar",
	Long:  "Sidecar for WebApp to provide Prometheus metrics",
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}
