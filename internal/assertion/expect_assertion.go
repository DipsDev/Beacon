package assertion

import (
	"fmt"
	"strings"
)

type Expect struct {
	ExpectedValue string
}

func (e *Expect) Assert(ctx *Context) error {
	actual := ctx.Output[*ctx.OutputIdx]

	if actual != e.ExpectedValue {
		return fmt.Errorf("expected '%s' but got '%s'", e.ExpectedValue, actual)
	}

	*ctx.OutputIdx++
	return nil
}

func newExpect(args []string) (Assertion, error) {
	return &Expect{
		ExpectedValue: strings.Join(args, " "),
	}, nil
}
