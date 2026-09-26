package reporter

import (
	"beacon/internal/assertion"
	"beacon/internal/parser"
)

type Writer interface {
	Write()
}

type Reporter struct {
	files []fileReport

	currentFile  *fileReport
	currentGroup *groupReport
}

func (c *Reporter) AddFile(file parser.File) {
	c.files = append(c.files, fileReport{file, nil})

	c.currentFile = &c.files[len(c.files)-1]
	c.currentGroup = nil
}

func (c *Reporter) AddGroup(group parser.Group) {
	c.currentFile.groups = append(c.currentFile.groups, groupReport{group, nil})

	c.currentGroup = &c.currentFile.groups[len(c.currentFile.groups)-1]
}

func (c *Reporter) AddAssertion(as assertion.Assertion, result error) {
	c.currentGroup.assertions = append(c.currentGroup.assertions, assertionReport{
		assertion: as,
		result:    result,
	})
}
