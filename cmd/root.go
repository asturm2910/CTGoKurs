package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CTGoKurs",
	Short: "Read Atom Feeds",
	Long: `Just a small CLI application. 
	Read Admon Feeds`,
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List news",
	Long:  "List first 5 Newsfeeds",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

var cmdDescribe = &cobra.Command{
	Use:   "describe [id]",
	Short: "Show details of an article",
	Long:  "Details for an article",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0])
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdDescribe)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
