package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version contains the current version.
	version = "dev"
	// BuildDate contains a string with the build date.
	buildDate = "unknown"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Display version and build information about WebApp Sidecar",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("webapp-sidecar %s\n", version)
		fmt.Printf("  Build date: %s\n", buildDate)
		fmt.Printf("  Built with: %s\n", runtime.Version())
	},
}
