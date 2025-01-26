package model

type Errs struct {
	error
}

func (err Errs) Message() string {
	return err.Error()
}
