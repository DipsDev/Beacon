package parser

import (
	"beacon/internal/assertion"
	"time"
)

type FileSettings struct {
	Timeout time.Duration
}

func defaultSettings() *FileSettings {
	return &FileSettings{
		Timeout: 5 * time.Second,
	}
}

type Group struct {
	Description string
	Line        int
	Assertions  []assertion.Assertion
}

func (g *Group) AppendAssertion(a assertion.Assertion) {
	g.Assertions = append(g.Assertions, a)
}

type File struct {
	Name     string
	Settings *FileSettings
	Groups   []Group
}
