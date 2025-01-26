package outerr

import (
	"errors"
)

type Errs struct {
	error
}

func (err Errs) Message() string {
	return err.Error()
}

func NewError(err any) Errs {
	return Errs{
		errors.New(log.PrettyString(err)),
	}
}

const (
	ContentTypeNotSet = "Content-Type header not set"
)
