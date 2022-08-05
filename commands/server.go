package commands

import (
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/parmcoder/website-checker-backend/controllers"
	"github.com/parmcoder/website-checker-backend/injection"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server command",
	Long:  "Run the server using Echo v4",
	Run: func(cmd *cobra.Command, args []string) {
		serverInitialize()
	},
}

func serverInitialize() {
	e := echo.New()
	server := injection.InitializeServer()

	handleServer(e, &server)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}

func handleServer(e *echo.Echo, server *controllers.Server) {
	e.POST("/text", (*server).CheckHealth)
	e.POST("/csv", (*server).CheckHealthCsv)
}
