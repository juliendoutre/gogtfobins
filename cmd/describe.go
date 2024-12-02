package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/juliendoutre/gogtfobins"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func describeCmd() *cobra.Command {
	var (
		function string
		format   string
	)

	cmd := &cobra.Command{
		Use:   "describe BINARY",
		Short: "Print out informations about a given binary.",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			index, err := gogtfobins.LoadIndex()
			if err != nil {
				return fmt.Errorf("building gtfobins index: %w", err)
			}

			object, err := extractObject(index, function, args[0])
			if err != nil {
				return err
			}

			content, err := formatObject(object, format)
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, content)

			return nil
		},
	}

	cmd.Flags().StringVar(
		&function, "function", "",
		"optional function to display information for",
	)

	cmd.Flags().StringVar(
		&format, "format", "yaml",
		"format of the output",
	)

	return cmd
}

func extractObject(index gogtfobins.Index, function, name string) (any, error) {
	gtfobin, ok := index[name]
	if !ok {
		return nil, ErrUnknownBinary
	}

	if function != "" {
		fun, ok := gtfobin.Functions[function]
		if !ok {
			return nil, ErrUnknownFunction
		}

		return fun, nil
	}

	return gtfobin.Functions, nil
}

func formatObject(object any, format string) (string, error) {
	switch format {
	case "json":
		content, err := json.MarshalIndent(object, "", "  ")
		if err != nil {
			return "", fmt.Errorf("encoding output to json: %w", err)
		}

		return string(content), nil
	case "yaml":
		content, err := yaml.Marshal(object)
		if err != nil {
			return "", fmt.Errorf("encoding output to yaml: %w", err)
		}

		return string(content), nil
	default:
		return "", ErrUnknownFormat
	}
}
