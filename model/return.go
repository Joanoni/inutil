package model

type Caio interface {
	// ReturnStruct[any] | ReturnStructError
	GetMessage() string
	GetData() any
	GetStatusCode() int
}

type ReturnStruct[T any] struct {
	Message    string `json:"message"`
	Data       T      `json:"data"`
	StatusCode int    `json:"-"`
}

func (rs ReturnStruct[T]) GetMessage() string {
	return rs.Message
}

func (rs ReturnStruct[T]) GetData() T {
	return rs.Data
}

func (rs ReturnStruct[T]) GetStatusCode() int {
	return rs.StatusCode
}

type ReturnStructError struct {
	Message    string `json:"message"`
	Error      *Errs  `json:"-"`
	StatusCode int    `json:"-"`
}

func (rse ReturnStructError) GetMessage() string {
	return rse.Message
}

func (rse ReturnStructError) GetData() any {
	return rse.Error
}

func (rse ReturnStructError) GetStatusCode() int {
	return rse.StatusCode
}

func (rse ReturnStructError) HasError() bool {
	return rse.Error != nil
}
