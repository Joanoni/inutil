package inutil

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
)

var server *Server_Model

func Server_Start(address string) *Server_Model {
	server = &Server_Model{}

	server.Router = pat.New()

	server.HTTPServer = &http.Server{
		Addr: address, //"0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      server.Router, // Pass our instance of gorilla/mux in.
	}

	server.middleware_ch = &middleware_context_model{
		Contexts: map[*http.Request]*Context{},
	}

	server.Router.Use(middleware_context_handler)
	server.Router.Use(middleware_log_handler)

	return server
}

func Oi() {
	Print("oi2")
}

func (s *Server_Model) Get(path string, h HandlerFunc) *mux.Route {
	return s.Router.Get(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Run() {
	LogF("Running server: %v", s.HTTPServer.Addr)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

func (s *Server_Model) Context(req *http.Request) *Context {
	return server.middleware_ch.Contexts[req]
}

func (s *Server_Model) JSON(c *Context, payload Return[any]) {
	payloadJ, err := json.Marshal(payload)
	if s.HandleError(c, err) {
		return
	}
	c.wr.Write(payloadJ)
	c.wr.WriteHeader(payload.Status)
	c.wr.Header().Set("Content-Type", ApplicationJSON)
}

func (s *Server_Model) HandleError(c *Context, err error) bool {
	if err != nil {
		LogF("HandleError: %v", err.Error())
		return true
	}
	return false
}
