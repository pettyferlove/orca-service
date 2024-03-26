package version

import (
	"fmt"
	"github.com/spf13/cobra"
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
	fmt.Println("1.0.0")
	return nil
}
