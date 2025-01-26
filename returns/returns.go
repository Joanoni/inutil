package return

// type ReturnStruct interface {
// 	ReturnStructT[any] | ReturnStructAny
// 	GetMessage() string
// 	GetData() any
// 	GetStatusCode() int
// }

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
	Error      *Error `json:"-"`
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

// type ReturnStructAny struct {
// 	Message    string `json:"message"`
// 	Data       any    `json:"data"`
// 	StatusCode int    `json:"-"`
// }

// func (rs *ReturnStructAny) GetMessage() string {
// 	return rs.Message
// }

// func (rs *ReturnStructAny) GetData() any {
// 	return rs.Data
// }

// func (rs *ReturnStructAny) GetStatusCode() int {
// 	return rs.StatusCode
// }

func ReturnBadRequest(err Error) ReturnStructError {
	return ReturnStructError{
		Message:    err.Error(),
		Error:      &err,
		StatusCode: StatusBadRequest,
	}
}

func ReturnInternalServerError(err Error) ReturnStructError {
	return ReturnStructError{
		Message:    err.Error(),
		Error:      &err,
		StatusCode: StatusInternalServerError,
	}
}

func ReturnOk[T any](data T) ReturnStruct[T] {
	return ReturnStruct[T]{
		Message:    "success",
		Data:       data,
		StatusCode: StatusOk,
	}
}

func ReturnEmpty[T any]() ReturnStruct[T] {
	return ReturnStruct[T]{}
}
