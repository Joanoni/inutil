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
	Payload any
	Header  http.Header
}

type RequestReponse[T any] struct {
	StatusCode int
	Body       T
}

var client *http.Client

func startRequest() {
	client = &http.Client{}
}

func Request[T any](input RequestInput, c *Context) (output ReturnStruct[T], outerr ReturnStructError) {
	defer PrintInternalFunction()()
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
					body, err = json.Marshal(input.Payload)
					if HandleError(err) {
						outerr = ReturnInternalServerError(ErrsFromError(err))
						return
					}
				}
			}
		}
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, bytes.NewReader(body))
		if HandleError(err) {
			outerr = ReturnInternalServerError(ErrsFromError(err))
			return
		}
	} else {
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, nil)
		if HandleError(err) {
			outerr = ReturnInternalServerError(ErrsFromError(err))
			return
		}
	}
	req.Header = input.Header
	resp, err := client.Do(req)
	if HandleError(err) {
		output = ReturnEmpty[T]()
		outerr = ReturnInternalServerError(ErrsFromError(err))
		return
	}
	var parsedBody T
	outerr = parseBody(resp.Body, resp.Header[HeaderContentType], &parsedBody)
	if outerr.HasError() {
		return
	}
	logInternal("parsedBody", PrettyString(parsedBody))
	output = ReturnOk(parsedBody)
	outerr = ReturnInternalServerError(ErrsFromError(err))
	return
}
