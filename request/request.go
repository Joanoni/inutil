package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

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
					body, err = json.Marshal(input.Payload.Body)
					if HandleError(err) {
						output = ReturnEmpty[T]()
						outerr = ReturnInternalServerError(NewError(err.Error()))
						return
					}
				}
			}
		}
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, bytes.NewReader(body))
		if HandleError(err) {
			output = ReturnEmpty[T]()
			outerr = ReturnInternalServerError(NewError(err.Error()))
			return
		}
	} else {
		req, err = http.NewRequest(strings.ToUpper(input.Method), input.Url, nil)
		if HandleError(err) {
			output = ReturnEmpty[T]()
			outerr = ReturnInternalServerError(NewError(err.Error()))
			return
		}
	}

	req.Header = input.Header

	resp, err := client.Do(req)
	if HandleError(err) {
		output = ReturnEmpty[T]()
		outerr = ReturnInternalServerError(NewError(err.Error()))
		return
	}

	var parsedBody T
	// outerr := parseBody(resp.Body, resp.Header[HeaderContentType], &parsedBody)

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(resp.Body)
	if HandleError(err) {
		output = ReturnEmpty[T]()
		outerr = ReturnInternalServerError(NewError(err.Error()))
		return
	}

	logInternal("Request: string(bodyBytes)", string(bodyBytes))

	if len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, &parsedBody)
		if HandleError(err) {
			output = ReturnEmpty[T]()
			outerr = ReturnInternalServerError(NewError(err.Error()))
			return
		}
	}

	logInternal("parsedBody", PrettyString(parsedBody))

	output = ReturnOk[T](parsedBody)
	outerr = ReturnInternalServerError(NewError(err.Error()))
	return
}
