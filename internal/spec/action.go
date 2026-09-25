package spec

import "errors"

type Action string

const (
	ActionExpect Action = "expect"
	ActionExit   Action = "exit"
)

func ParseAction(action string) (Action, error) {
	switch action {
	case "expect":
		return ActionExpect, nil
	case "exit":
		return ActionExit, nil
	default:
		return "", errors.New("invalid action")
	}
}
