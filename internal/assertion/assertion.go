package assertion

import "fmt"

type Context struct {
	Output     []string
	OutputIdx  *int
	ExitStatus int
}

type Assertion interface {
	Assert(ctx *Context) error
}

type assertionCreator func(args []string) (Assertion, error)

var assertionsRegistry = map[string]assertionCreator{}

func registerAssertion(name string, fn assertionCreator) {
	assertionsRegistry[name] = fn
}

func Parse(action string, args []string) (Assertion, error) {
	if ac, ok := assertionsRegistry[action]; ok {
		return ac(args)
	}

	return nil, fmt.Errorf("invalid assertion action: %s", action)
}

func init() {
	registerAssertion("expect", newExpect)
	registerAssertion("exit", newExit)
}
