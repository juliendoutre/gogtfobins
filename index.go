package gogtfobins

import (
	"bytes"
	"embed"
	"fmt"
	"path"
	"strings"

	set "github.com/deckarep/golang-set/v2"
	"gopkg.in/yaml.v3"
)

// Index lists available Functions for each GTFOBin.
type Index map[string]GTFOBin

type GTFOBin struct {
	Functions map[string][]Function `yaml:"functions"`
}

type Function struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}

// ReverseIndex lists GTBOBins providing each Function.
type ReverseIndex map[string]set.Set[string]

//go:embed GTFOBins.github.io/_gtfobins/*.md
var markdowns embed.FS

const rootDir = "GTFOBins.github.io/_gtfobins"

func LoadIndex() (Index, error) {
	dir, err := markdowns.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("reading gtfobins embedded directory: %w", err)
	}

	index := Index{}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		content, err := markdowns.ReadFile(path.Join(rootDir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("reading gtfobins embedded file: %w", err)
		}

		var gtfobin GTFOBin
		if err := yaml.Unmarshal(bytes.Trim(content, "-"), &gtfobin); err != nil {
			return nil, fmt.Errorf("parsing gtfobins embedded file: %w", err)
		}

		name := strings.TrimSuffix(file.Name(), ".md")

		index[name] = gtfobin
	}

	return index, nil
}

func BuildReverseIndex(index Index) ReverseIndex {
	reverseIndex := ReverseIndex{}

	for gtfobinName, gtfobin := range index {
		for function := range gtfobin.Functions {
			if entry, ok := reverseIndex[function]; ok {
				entry.Add(gtfobinName)
			} else {
				reverseIndex[function] = set.NewSet(gtfobinName)
			}
		}
	}

	return reverseIndex
}
