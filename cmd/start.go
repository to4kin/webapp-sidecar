package cmd

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/to4kin/webapp-sidecar/internal/app/metricserver"
)

var (
	configPath string
)

func init() {
	startCmd.PersistentFlags().StringVarP(&configPath, "config-path", "c", "configs/webappsidecar.toml", "path to config file")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start WebApp Sidecar",
	Long: `Start WebApp Sidecar with config file
Simply execute webapp-sidecar start -c path/to/config/file.toml
or skip this flag to use default path`,
	Run: func(cmd *cobra.Command, args []string) {
		config := metricserver.NewConfig()
		if _, err := toml.DecodeFile(configPath, config); err != nil {
			logrus.Fatal(err)
		}

		if err := metricserver.Start(config); err != nil {
			logrus.Fatal(err)
		}
	},
}
