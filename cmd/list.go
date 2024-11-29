package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	set "github.com/deckarep/golang-set/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/juliendoutre/gogtfobins"
	"github.com/spf13/cobra"
)

var function string //nolint: gochecknoglobals

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available binaries with their possible functions.",
		RunE: func(_ *cobra.Command, _ []string) error {
			index, err := gogtfobins.LoadIndex()
			if err != nil {
				return fmt.Errorf("building gtfobins index: %w", err)
			}

			candidates, err := listCandidateBinaries(function, index)
			if err != nil {
				return err
			}

			binaries, err := listAvailableBinaries(candidates)
			if err != nil {
				return err
			}

			tableWriter := table.NewWriter()
			tableWriter.SetStyle(table.StyleColoredBright)
			tableWriter.AppendHeader(table.Row{"name", "path", "functions"})

			for binary, path := range binaries {
				tableWriter.AppendRow(table.Row{
					binary,
					path,
					strings.Join(
						set.NewSetFromMapKeys(index[binary].Functions).ToSlice(),
						",",
					),
				})
			}

			fmt.Fprintln(os.Stdout, tableWriter.Render())

			return nil
		},
	}

	cmd.Flags().StringVar(
		&function, "function", "",
		"optional function to filter binaries by",
	)

	return cmd
}

func listCandidateBinaries(function string, index gogtfobins.Index) (set.Set[string], error) {
	// In case a specific function is requested, let's leverage the reverse index to speed up the research.
	if function != "" {
		reverseIndex := gogtfobins.BuildReverseIndex(index)

		potentialBinaries, ok := reverseIndex[function]
		if !ok {
			return nil, ErrUnknwonFunction
		}

		return potentialBinaries, nil
	}

	// Else just check for all binaries
	return set.NewSetFromMapKeys(index), nil
}

func listAvailableBinaries(candidates set.Set[string]) (map[string]string, error) {
	availableBinaries := map[string]string{}

	for candidate := range candidates.Iter() {
		binaryPath, err := exec.LookPath(candidate)
		if err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				continue
			}

			return nil, fmt.Errorf("checking for binary: %w", err)
		}

		if binaryPath != "" {
			availableBinaries[candidate] = binaryPath
		}
	}

	return availableBinaries, nil
}
