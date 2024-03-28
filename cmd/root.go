package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "tt",
	Short: "A supplementary tool for github.com/alwqx/translate.",
	Long: `tt is a CLI for generating template Markdown post.
It is developed to quickly generate template post for https://github.com/alwqx/translate repo.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
