package inutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type RequestInput struct {
	Method  string
	Url     string
	Payload *RequestPayloadInput
}

type RequestPayloadInput struct {
	Body        any
	ContentType string
}

type RequestReponse[T any] struct {
	StatusCode int
	Body       T
}

var client *http.Client

func startRequest() {
	client = &http.Client{}
}

func Request[T any](input RequestInput, c *Context) Return[RequestReponse[*T]] {
	headers := http.Header{}
	var err error
	var body []byte
	var req *http.Request

	if input.Payload != nil {
		switch input.Payload.ContentType {
		case ApplicationJSON:
			headers.Set(HeaderContentType, ApplicationJSON)
			body, err = json.Marshal(input.Payload.Body)
			if c.HandleError(err) {
				return Return[RequestReponse[*T]]{
					Message:    err.Error(),
					Data:       nil,
					Success:    false,
					StatusCode: StatusBadRequest,
				}
			}
		}
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, bytes.NewReader(body))
		if c.HandleError(err) {
			return Return[RequestReponse[*T]]{
				Message:    err.Error(),
				Data:       nil,
				Success:    false,
				StatusCode: StatusBadRequest,
			}
		}
	} else {
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, nil)
		if c.HandleError(err) {
			return Return[RequestReponse[*T]]{
				Message:    err.Error(),
				Data:       nil,
				Success:    false,
				StatusCode: StatusBadRequest,
			}
		}
	}

	resp, err := client.Do(req)
	if c.HandleError(err) {
		return Return[RequestReponse[*T]]{
			Message:    err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: StatusBadRequest,
		}
	}

	var bodyBytes []byte
	_, err = resp.Body.Read(bodyBytes)
	if c.HandleError(err) {
		return Return[RequestReponse[*T]]{
			Message: err.Error(),
			Data: &RequestReponse[*T]{
				StatusCode: resp.StatusCode,
			},
			Success:    false,
			StatusCode: StatusBadRequest,
		}
	}

	var parsedBody *T
	if len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, parsedBody)
		if c.HandleError(err) {
			return Return[RequestReponse[*T]]{
				Message: err.Error(),
				Data: &RequestReponse[*T]{
					StatusCode: resp.StatusCode,
				},
				Success:    false,
				StatusCode: StatusBadRequest,
			}
		}
	}

	return Return[RequestReponse[*T]]{
		Message: "success",
		Data: &RequestReponse[*T]{
			StatusCode: resp.StatusCode,
			Body:       parsedBody,
		},
		Success:    true,
		StatusCode: StatusOK,
	}
}
