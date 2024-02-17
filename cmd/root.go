package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Dymplex CLI",
	Short: "Root command",
	Long:  "This is the root command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root executed once")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
