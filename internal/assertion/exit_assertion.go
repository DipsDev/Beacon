package assertion

import (
	"fmt"
	"strconv"
)

type Exit struct {
	ExpectedValue int
}

func (e *Exit) Assert(ctx *Context) error {
	if ctx.ExitStatus != e.ExpectedValue {
		return fmt.Errorf("expected status code %d but got %d", e.ExpectedValue, ctx.ExitStatus)
	}

	*ctx.OutputIdx++
	return nil
}

func newExit(args []string) (Assertion, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("exit expected 1 argument, got %d", len(args))
	}

	expectedValue, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("exit expected integer but got '%s'", args[0])
	}

	return &Exit{
		ExpectedValue: expectedValue,
	}, nil
}
