package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server command",
	Long:  "Run the server using Echo v4",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logrus.Info("Running server")
	},
}
