package runner

import (
	"beacon/internal/assertion"
	"beacon/internal/parser"
	"beacon/internal/reporter"
	"fmt"
	"os"
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

	c := reporter.NewConsole()

	for _, testFile := range testFiles {
		file, err := parser.ParseTestFile(testFile, cfg.CommentPrefix)
		if err != nil {
			fmt.Printf("error parsing test file: %s\n, %v", testFile, err)
			continue
		}

		cmd, err := executeFile(cfg.ExecutablePath, []string{testFile}, *file.Settings)
		if err != nil {
			fmt.Printf("error executing test file: %v\n", err)
			continue
		}

		cleanedOutput := strings.ReplaceAll(string(cmd.Output), "\r\n", "\n")
		actualOutput := strings.Split(strings.TrimSpace(cleanedOutput), "\n")

		outputIndex := 0
		ctx := assertion.Context{
			Output:    actualOutput,
			OutputIdx: &outputIndex,
		}

		c.AddFile(*file)
		for _, group := range file.Groups {

			c.AddGroup(group)
			for _, groupAssertion := range group.Assertions {
				err = groupAssertion.Assert(&ctx)
				c.AddAssertion(groupAssertion, err)
			}

		}

	}
	c.Write()

	return 0
}
