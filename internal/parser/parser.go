package parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const Prefix = "beacon:"

func ParseTestFile(path string, commentPrefix string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	targetPrefix := commentPrefix + Prefix

	parseState := ParseState{File: &File{
		Name:     filepath.Base(path),
		Settings: defaultSettings(),
	}}

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, targetPrefix) {
			continue
		}

		cleaned := strings.TrimSpace(strings.TrimPrefix(line, targetPrefix))
		parts := strings.Fields(cleaned)

		if len(parts) == 0 {
			return nil, fmt.Errorf("%s:%d empty beacon directive", path, lineNumber)
		}

		directiveName := parts[0]
		directiveArgs := parts[1:]

		directiveParser, ok := directiveParsers[directiveName]
		if !ok {
			return nil, fmt.Errorf("%s:%d directive '%s' not recognized", path, lineNumber, directiveName)
		}

		directive, err := directiveParser(directiveArgs)
		if err != nil {
			return nil, err
		}

		err = directive.Apply(&parseState)
		if err != nil {
			return nil, err
		}

	}
	return parseState.File, scanner.Err()
}
