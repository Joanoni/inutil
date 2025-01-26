package inutil

// type Return[T any] interface {
// 	ReturnStruct[T] | ReturnStructAny
// 	GetMessage() string
// 	GetData() *T
// 	GetStatusCode() int
// }

type ReturnStructT[T any] struct {
	Message    string `json:"message"`
	Data       *T     `json:"data"`
	StatusCode int    `json:"-"`
}

type ReturnStruct struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"-"`
}

func (rs *ReturnStructT[T]) GetMessage() string {
	return rs.Message
}

func (rs *ReturnStructT[T]) GetData() *T {
	return rs.Data
}

func (rs *ReturnStructT[T]) GetStatusCode() int {
	return rs.StatusCode
}

func (rs *ReturnStruct) GetMessage() string {
	return rs.Message
}

func (rs *ReturnStruct) GetData() any {
	return rs.Data
}

func (rs *ReturnStruct) GetStatusCode() int {
	return rs.StatusCode
}

func ReturnBadRequest(format string, values ...any) ReturnStruct {
	return ReturnStruct{
		Message:    SprintF(format, values),
		Data:       nil,
		StatusCode: StatusBadRequest,
	}
}

func ReturnInternalServerError(format string, values ...any) ReturnStruct {
	return ReturnStruct{
		Message:    SprintF(format, values),
		Data:       nil,
		StatusCode: StatusInternalServerError,
	}
}

func ReturnOk(format string, values ...any) ReturnStruct {
	return ReturnStruct{
		Message:    SprintF(format, values),
		Data:       nil,
		StatusCode: StatusOk,
	}
}
