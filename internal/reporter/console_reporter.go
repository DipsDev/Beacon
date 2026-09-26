package reporter

import (
	"beacon/internal/assertion"
	"errors"
	"fmt"
	"strings"
)

type Console struct {
	Reporter
}

func NewConsole() *Console {
	return &Console{}
}

func printError(err error) {
	switch err.(type) {
	case *assertion.ErrUnmatched:
		{
			var erru *assertion.ErrUnmatched
			errors.As(err, &erru)
			fmt.Printf("      Expected: %v\n      Got: %v\n\n", erru.Expected, erru.Actual)
		}
	default:
		fmt.Printf("      Unhandled: %v\n", err)
	}
}

func (c *Console) Write() {
	fmt.Println("running beacon tests...")
	fmt.Println(strings.Repeat("-", 40))

	okFiles := 0

	okGroups := 0
	totalGroups := 0

	okAssertions := 0
	totalAssertions := 0

	for _, file := range c.files {
		totalGroups += file.groupCount()
		totalAssertions += file.assertionCount()

		if file.ok() {
			okFiles++

			okGroups += file.groupCount()
			okAssertions += file.assertionCount()

			fmt.Printf("✓ %s (%d)\n", file.file.Name, file.assertionCount())
			continue
		}

		fmt.Printf("✕ %s (%d)\n", file.file.Name, file.assertionCount())

		for _, grp := range file.groups {
			if !grp.ok() {
				for _, asser := range grp.assertions {
					if !asser.ok() {
						fmt.Printf("   └─ %s > %s\n", grp.group.Description, asser.assertion)
						printError(asser.result)
					} else {
						okAssertions++
					}
				}

			} else {
				okGroups++
			}
		}

		fmt.Println()
	}

	// Final Summary Block
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Test Files:  %d passed, %d total\n", okFiles, len(c.files))
	fmt.Printf("Test Suites: %d passed, %d total\n", okGroups, totalGroups)
	fmt.Printf("Tests:       %d failed, %d passed, %d total\n", totalAssertions-okAssertions, okAssertions, totalAssertions)
	fmt.Println(strings.Repeat("=", 50))
}
