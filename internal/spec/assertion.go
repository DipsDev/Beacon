package spec

import (
	"fmt"
	"strconv"
)

type Assertion struct {
	Action        Action
	ExpectedValue string
	Line          int
}

type Context struct {
	Output   []string
	ExitCode int
	Index    int
}

type AssertResultType string

const (
	AssertResultSuccess AssertResultType = "passed"
	AssertResultFailure AssertResultType = "failed"
	AssertResultSkipped AssertResultType = "skipped"
)

type AssertResult struct {
	Type AssertResultType
	Err  error
}

func resultFromError(err error) AssertResult {
	return AssertResult{
		Type: AssertResultFailure,
		Err:  err,
	}
}

func NewAssertion(action Action, expected string, line int) *Assertion {
	return &Assertion{
		Action:        action,
		ExpectedValue: expected,
		Line:          line,
	}
}

func (a *Assertion) expect(ctx Context) AssertResult {
	if ctx.Output[ctx.Index] != a.ExpectedValue {
		return resultFromError(fmt.Errorf("expected %s but got %s", a.ExpectedValue, ctx.Output))
	}

	return AssertResult{Type: AssertResultSuccess}
}

func (a *Assertion) exit(ctx Context) AssertResult {
	value, err := strconv.Atoi(a.ExpectedValue)
	if err != nil {
		return resultFromError(fmt.Errorf("exit expected a number but got %s", strconv.Quote(a.ExpectedValue)))
	}

	if value != ctx.ExitCode {
		return resultFromError(fmt.Errorf("expected exit code %s but got %d", a.ExpectedValue, ctx.ExitCode))
	}

	return AssertResult{Type: AssertResultSuccess}
}

func (a *Assertion) Test(ctx Context) AssertResult {
	switch a.Action {
	case ActionExpect:
		return a.expect(ctx)
	case ActionExit:
		return a.exit(ctx)
	default:
		return AssertResult{
			Type: AssertResultFailure,
			Err:  fmt.Errorf("unknown action %s", a.Action),
		}
	}
}
