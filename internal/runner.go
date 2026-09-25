package internal

import (
	"beacon/internal/spec"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type Config struct {
	ExecutablePath string
	CommentPrefix  string
	TestDirPath    string
}

func RunTests(cfg Config) int {
	if _, err := os.Stat(cfg.ExecutablePath); os.IsNotExist(err) {
		fmt.Printf("Error: executable not found at %s.\n", cfg.ExecutablePath)
		return 1
	}

	pattern := filepath.Join(cfg.TestDirPath, "*.beacon")
	testFiles, err := filepath.Glob(pattern)

	if err != nil || len(testFiles) == 0 {
		fmt.Printf("No test files found matching pattern: %s\n", pattern)
		return 0
	}

	sort.Strings(testFiles)

	fmt.Println("running beacon tests...\n" + strings.Repeat("-", 40))

	assertionsFailed := 0
	totalAssertions := 0

	for _, testFile := range testFiles {
		// Inside your test execution loop:
		assertions, err := ParseTestFile(testFile, cfg.CommentPrefix)
		if err != nil || len(assertions) == 0 {
			fmt.Printf("skipping: %s (no beacons assertions found)\n", filepath.Base(testFile))
			continue
		}

		cmd := exec.Command(cfg.ExecutablePath, testFile)
		outputBytes, err := cmd.CombinedOutput()

		cleanedOutput := strings.ReplaceAll(string(outputBytes), "\r\n", "\n")
		actualOutput := strings.Split(strings.TrimSpace(cleanedOutput), "\n")

		totalAssertions += len(assertions)

		for i, assertion := range assertions {
			fmt.Printf("<%s:%d> %s %s ... ", filepath.Base(testFile), assertion.Line, assertion.Action, assertion.ExpectedValue)
			outputContext := spec.Context{
				Output:   actualOutput,
				ExitCode: cmd.ProcessState.ExitCode(),
				Index:    i,
			}

			result := assertion.Test(outputContext)
			if result.Err != nil {
				fmt.Printf("[%s] (%s)\n", result.Type, result.Err)
				assertionsFailed++
				continue
			}
			fmt.Printf("[%s]\n", result.Type)

		}
		fmt.Println()
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("finished running test suites. assertions failed %d/%d", assertionsFailed, totalAssertions)

	return 0
}
