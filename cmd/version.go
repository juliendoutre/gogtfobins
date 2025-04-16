package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// This is set at build time.
var (
	version = "unknown"
	commit  = "unknown" //nolint:gochecknoglobals
	date    = "unknown" //nolint:gochecknoglobals
)

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print out the CLI version.",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Fprintf(os.Stdout, "gogtfobins v%s, commit %s, built at %s\n", version, commit, date)
		},
	}

	return cmd
}
