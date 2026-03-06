package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ship-htmx",
	Short: "Ship efficient HTMX websites and applications",
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var cfgfile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgfile, "config", "", "config file (default: $HOME/.ship-htmx.yaml)")
}
