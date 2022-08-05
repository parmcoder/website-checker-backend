package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Execute This function serve as the reader of command in terminal
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root command",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serverCmd)
}

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logrus.Info(os.Getenv("APP_ENV"))
	logrus.Info(os.Getenv("PORT"))
}
