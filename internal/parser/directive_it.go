package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type itDirective struct {
	description string
}

func (d *itDirective) Apply(state *ParseState) error {
	state.newGroup().Description = d.description
	return nil
}

func newItDirective(args []string) (Directive, error) {
	joinedArgs := strings.Join(args, " ")
	desc, err := strconv.Unquote(joinedArgs)
	if err != nil {
		return nil, fmt.Errorf("it description should be quoted, got %s", joinedArgs)
	}

	return &itDirective{
		description: desc,
	}, nil
}
