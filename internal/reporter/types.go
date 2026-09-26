package reporter

import (
	"beacon/internal/assertion"
	"beacon/internal/parser"
)

type fileReport struct {
	file   parser.File
	groups []groupReport
}

type groupReport struct {
	group      parser.Group
	assertions []assertionReport
}

type assertionReport struct {
	assertion assertion.Assertion
	result    error
}

func (ar *assertionReport) ok() bool {
	return ar.result == nil
}

func (gr *groupReport) ok() bool {
	for _, ar := range gr.assertions {
		if !ar.ok() {
			return false
		}
	}
	return true
}

func (gr *groupReport) assertionCount() int {
	return len(gr.assertions)
}

func (fr *fileReport) ok() bool {
	for _, gr := range fr.groups {
		if !gr.ok() {
			return false
		}
	}

	return true
}

func (fr *fileReport) assertionCount() int {
	sum := 0
	for _, gr := range fr.groups {
		sum += gr.assertionCount()
	}
	return sum
}

func (fr *fileReport) groupCount() int {
	return len(fr.groups)
}
