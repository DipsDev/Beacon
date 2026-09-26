package parser

import (
	"fmt"
)

type ParseState struct {
	File *File
}

func (s *ParseState) currentGroup() *Group {
	if len(s.File.Groups) == 0 {
		s.File.Groups = append(s.File.Groups, Group{})
	}
	return &s.File.Groups[len(s.File.Groups)-1]
}

type Directive interface {
	Apply(state *ParseState) error
}

type DirectiveParser func(args []string) (Directive, error)

var directiveParsers = map[string]DirectiveParser{}

func registerDirective(keyword string, parser DirectiveParser) {
	if _, exists := directiveParsers[keyword]; exists {
		panic(fmt.Sprintf("beacon: directive %q registered twice", keyword))
	}
	directiveParsers[keyword] = parser
}

func init() {
	registerAssertionDirective("expect")
	registerAssertionDirective("exit")
}
