package assertion

import "fmt"

type ErrUnmatched struct {
	Expected interface{}
	Actual   interface{}
}

func (e *ErrUnmatched) Error() string {
	return fmt.Sprintf("expected %v, got %v", e.Expected, e.Actual)
}
