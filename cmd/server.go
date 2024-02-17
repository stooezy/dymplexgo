package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stooezy/dymplex/internal/api"
	"github.com/stooezy/dymplex/internal/api/router"
	"github.com/stooezy/dymplex/internal/config"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	config := config.DefaultServerConfig()

	s := api.NewServer(config)
	router.Init(s)
	s.Start()
}
