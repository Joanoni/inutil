package model

import "net/http"

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
