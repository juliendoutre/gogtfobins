package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	set "github.com/deckarep/golang-set/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/juliendoutre/gogtfobins"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func listCmd() *cobra.Command {
	var (
		function string
		format   string
	)

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

			entries := buildListEntries(index, binaries)

			content, err := formatList(entries, format)
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, content)

			return nil
		},
	}

	cmd.Flags().StringVar(
		&function, "function", "",
		"optional function to filter binaries by",
	)

	cmd.Flags().StringVar(
		&format, "format", "table",
		"format of the output",
	)

	return cmd
}

func listCandidateBinaries(function string, index gogtfobins.Index) (set.Set[string], error) {
	// In case a specific function is requested, let's leverage the reverse index to speed up the research.
	if function != "" {
		reverseIndex := gogtfobins.BuildReverseIndex(index)

		potentialBinaries, ok := reverseIndex[function]
		if !ok {
			return nil, ErrUnknownFunction
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

type listEntry struct {
	Name      string   `json:"name"      yaml:"name"`
	Path      string   `json:"path"      yaml:"path"`
	Functions []string `json:"functions" yaml:"functions"`
}

func buildListEntries(index gogtfobins.Index, binaries map[string]string) []listEntry {
	entries := []listEntry{}

	for binary, path := range binaries {
		entries = append(entries, listEntry{
			Name:      binary,
			Path:      path,
			Functions: set.NewSetFromMapKeys(index[binary].Functions).ToSlice(),
		})
	}

	return entries
}

func formatList(entries []listEntry, format string) (string, error) {
	switch format {
	case "yaml":
		content, err := yaml.Marshal(entries)
		if err != nil {
			return "", fmt.Errorf("encoding output to yaml: %w", err)
		}

		return string(content), nil
	case "json":
		content, err := json.MarshalIndent(entries, "", "  ")
		if err != nil {
			return "", fmt.Errorf("encoding output to json: %w", err)
		}

		return string(content), nil
	case "table":
		tableWriter := table.NewWriter()
		tableWriter.SetStyle(table.StyleColoredBright)
		tableWriter.AppendHeader(table.Row{"name", "path", "functions"})

		for _, entry := range entries {
			tableWriter.AppendRow(table.Row{
				entry.Name,
				entry.Path,
				strings.Join(entry.Functions, ","),
			})
		}

		return tableWriter.Render(), nil
	default:
		return "", ErrUnknownFormat
	}
}
