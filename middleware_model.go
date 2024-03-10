package inutil

import "net/http"

type middleware_context_model struct {
	Contexts map[*http.Request]*Context
}

type Context struct {
	wr   http.ResponseWriter
	req  *http.Request
	data map[string]any
	err  error
}

const ()
