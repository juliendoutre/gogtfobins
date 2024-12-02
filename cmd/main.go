package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "gogtfobins",
		Short:        "List Linux binaries potentially exploitable if misconfigured.",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		versionCmd(),
		listCmd(),
		describeCmd(),
		exploitCmd(),
	)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
