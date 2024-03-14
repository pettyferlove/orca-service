package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"orca-service/command/api"
	"orca-service/command/version"
	"os"
)

var root = &cobra.Command{
	Use:          "orca",
	Short:        "orca",
	SilenceUsage: true,
	Long:         "orca",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("parameter exception")
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	root.AddCommand(version.Command)
	root.AddCommand(api.Command)
}

func tip() {
	message := "Welcome to use orca, you can use -h to view the command"
	fmt.Printf("%s\n", message)
}

func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(-1)
	}
}
