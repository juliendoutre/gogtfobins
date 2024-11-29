package main

import (
	"fmt"
	"os"

	"github.com/juliendoutre/gogtfobins"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func describeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe BINARY",
		Short: "Print out informations about a given binary.",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			index, err := gogtfobins.LoadIndex()
			if err != nil {
				return fmt.Errorf("building gtfobins index: %w", err)
			}

			gtfobin, ok := index[args[0]]
			if !ok {
				return ErrUnknwonBinary
			}

			content, err := yaml.Marshal(gtfobin.Functions)
			if err != nil {
				return fmt.Errorf("encoding a description to yaml: %w", err)
			}

			fmt.Fprintln(os.Stdout, string(content))

			return nil
		},
	}

	return cmd
}
