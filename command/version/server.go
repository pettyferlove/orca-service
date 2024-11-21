package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version string
	Build   string
)

var (
	Command = &cobra.Command{
		Use:     "version",
		Short:   "Show version",
		Example: "orca version",
		PreRun: func(command *cobra.Command, args []string) {

		},
		RunE: func(command *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println("Version: ", Version)
	fmt.Println("Build: ", Build)
	return nil
}
