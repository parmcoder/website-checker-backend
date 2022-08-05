package commands

import (
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
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
		serverInitialize()
	},
}

func serverInitialize() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}
