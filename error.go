package gd

import (
	"errors"
)

func NewErr(str string) error {
	return errors.New(str)
}
