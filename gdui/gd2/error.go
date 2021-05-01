package gd2

import (
	"errors"
)

func NewErr(str string) error {
	return errors.New(str)
}
