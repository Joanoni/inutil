package inutil

import (
	"net/http"

	"github.com/gorilla/pat"
)

type HandlerFunc func(*Context)

type Return[V any] struct {
	Message string `json:"message"`
	Data    V      `json:"data"`
	Success bool   `json:"success"`
	Status  int    `json:"-"`
}

type Server_Model struct {
	HTTPServer    *http.Server
	Router        *pat.Router
	middleware_ch *middleware_context_model
}
