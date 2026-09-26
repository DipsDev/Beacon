package assertion

import (
	"fmt"
	"strings"
)

type Expect struct {
	ExpectedValue string
}

func (e *Expect) Assert(ctx *Context) error {
	if *ctx.OutputIdx >= len(ctx.Output) {
		return fmt.Errorf("too much outputs")

	}

	actual := ctx.Output[*ctx.OutputIdx]
	*ctx.OutputIdx++

	if actual != e.ExpectedValue {
		return &ErrUnmatched{
			Expected: e.ExpectedValue,
			Actual:   actual,
		}
	}

	return nil
}

func (e *Expect) String() string {
	return fmt.Sprintf("expect %s", e.ExpectedValue)
}

func newExpect(args []string) (Assertion, error) {
	return &Expect{
		ExpectedValue: strings.Join(args, " "),
	}, nil
}
