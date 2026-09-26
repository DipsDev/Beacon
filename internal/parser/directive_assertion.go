package parser

import "beacon/internal/assertion"

type assertionDirective struct {
	action string
	args   []string
}

func (a *assertionDirective) Apply(state *ParseState) error {
	parsedAssertion, err := assertion.Parse(a.action, a.args)
	if err != nil {
		return err
	}

	state.currentGroup().AppendAssertion(parsedAssertion)
	return nil
}

func registerAssertionDirective(action string) {
	registerDirective(action,
		func(args []string) (Directive, error) {
			return &assertionDirective{action: action, args: args}, nil
		})
}
