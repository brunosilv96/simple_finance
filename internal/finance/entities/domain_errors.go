package finance

import "fmt"

type InputCannotBeNil struct {
	Input string
}

func (error *InputCannotBeNil) Error() string {
	return fmt.Sprintf("the input %s cannot be nil or empty", error.Input)
}