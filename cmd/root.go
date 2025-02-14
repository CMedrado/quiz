package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const apiBaseURL = "http://localhost:5001"

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "CLI para o quiz interativo",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}
