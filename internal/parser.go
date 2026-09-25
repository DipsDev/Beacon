package internal

import (
	"beacon/internal/spec"
	"bufio"
	"os"
	"strings"
)

const Prefix = "beacon:"

func ParseTestFile(path string, commentPrefix string) ([]spec.Assertion, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	targetPrefix := commentPrefix + Prefix
	var assertions []spec.Assertion

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, targetPrefix) {
			continue
		}

		cleaned := strings.TrimPrefix(line, targetPrefix)
		parts := strings.Fields(cleaned)

		action := parts[0]
		actionArgs := parts[1]

		convAction, err := spec.ParseAction(action)
		if err != nil {
			return nil, err
		}

		assertions = append(assertions, *spec.NewAssertion(
			convAction,
			actionArgs,
			lineNumber,
		))

	}
	return assertions, scanner.Err()
}
