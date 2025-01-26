package inutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type RequestInput struct {
	Method  string
	Url     string
	Payload *RequestPayloadInput
	Header  http.Header
}

type RequestPayloadInput struct {
	Body any
}

type RequestReponse[T any] struct {
	StatusCode int
	Body       T
}

var client *http.Client

func startRequest() {
	client = &http.Client{}
}

func Request[T any](input RequestInput, c *Context) ReturnStructT[RequestReponse[*T]] {
	var err error
	var req *http.Request

	if input.Header == nil {
		input.Header = http.Header{}
	}

	logInternal("input", PrettyString(input))

	if input.Payload != nil {
		var body []byte
		for name, value := range input.Header {
			if name == HeaderContentType && len(value) > 0 {
				switch value[0] {
				case ApplicationJSON:
					body, err = json.Marshal(input.Payload.Body)
					if c.HandleError(err) {
						return ReturnStructT[RequestReponse[*T]]{
							Message:    err.Error(),
							Data:       nil,
							StatusCode: StatusBadRequest,
						}
					}
				}
			}
		}
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, bytes.NewReader(body))
		if c.HandleError(err) {
			return ReturnStructT[RequestReponse[*T]]{
				Message:    err.Error(),
				Data:       nil,
				StatusCode: StatusBadRequest,
			}
		}
	} else {
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, nil)
		if c.HandleError(err) {
			return ReturnStructT[RequestReponse[*T]]{
				Message:    err.Error(),
				Data:       nil,
				StatusCode: StatusBadRequest,
			}
		}
	}

	req.Header = input.Header

	resp, err := client.Do(req)
	if c.HandleError(err) {
		return ReturnStructT[RequestReponse[*T]]{
			Message:    err.Error(),
			Data:       nil,
			StatusCode: StatusBadRequest,
		}
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(resp.Body)
	if c.HandleError(err) {
		return ReturnStructT[RequestReponse[*T]]{
			Message: err.Error(),
			Data: &RequestReponse[*T]{
				StatusCode: resp.StatusCode,
			},
			StatusCode: StatusBadRequest,
		}
	}

	logInternal("Request: string(bodyBytes)", string(bodyBytes))

	var parsedBody T
	if len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, &parsedBody)
		if c.HandleError(err) {
			return ReturnStructT[RequestReponse[*T]]{
				Message: err.Error(),
				Data: &RequestReponse[*T]{
					StatusCode: resp.StatusCode,
				},
				StatusCode: StatusBadRequest,
			}
		}
	}

	logInternal("parsedBody", PrettyString(parsedBody))

	return ReturnStructT[RequestReponse[*T]]{
		Message: "success",
		Data: &RequestReponse[*T]{
			StatusCode: resp.StatusCode,
			Body:       &parsedBody,
		},
		StatusCode: StatusOk,
	}
}
