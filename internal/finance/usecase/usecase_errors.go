package usecase

import (
	"errors"
	"fmt"
)

var (
	CategoryNotFound = errors.New("category not found")
)

type InputCannotBeNil struct {
	Input string
}

func (error *InputCannotBeNil) Error() string {
	return fmt.Sprintf("the input %s cannot be nil or empty", error.Input)
}