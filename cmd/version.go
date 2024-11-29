package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// This is set at build time.
var Version = "unknown" //nolint:gochecknoglobals

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print out the CLI version.",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Fprintln(os.Stdout, "gogtfobins version: "+Version)
		},
	}

	return cmd
}
