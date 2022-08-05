package commands

import (
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/parmcoder/website-checker-backend/controllers"
	"github.com/parmcoder/website-checker-backend/injection"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server command",
	Long:  "Run the server using Echo v4",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logrus.Info("Running server")
		serverInitialize()
	},
}

func serverInitialize() {
	e := echo.New()
	server := injection.InitializeServer()

	e.POST("/", server.CheckHealth)

	e.Logger.Fatal(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}

func handleServer(e echo.Context, server controllers.Server) {
}
