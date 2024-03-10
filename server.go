package inutil

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
)

var server *Server_Model

func server_Start(ss *Start_Server) *Server_Model {
	server = &Server_Model{}

	server.Router = pat.New()

	server.HTTPServer = &http.Server{
		Addr: ss.Address, //"0.0.0.0:8080",
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

func (s *Server_Model) Get(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodGet, path)
	return s.Router.Get(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Head(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodHead, path)
	return s.Router.Head(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Post(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodPost, path)
	return s.Router.Post(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Put(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodPut, path)
	return s.Router.Put(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Patch(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodPatch, path)
	return s.Router.Patch(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Delete(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodDelete, path)
	return s.Router.Delete(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Options(path string, h HandlerFunc) *mux.Route {
	internalLogF("Method: %v, Path: %v", MethodOptions, path)
	return s.Router.Options(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Run() {
	internalLogF("Running server: %v", s.HTTPServer.Addr)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

func (s *Server_Model) Context(req *http.Request) *Context {
	return server.middleware_ch.Contexts[req]
}

func (c *Context) JSON(payload Return[any]) {
	payloadJ, err := json.Marshal(payload)
	if c.HandleError(err) {
		return
	}
	c.wr.WriteHeader(payload.Status)
	c.wr.Write(payloadJ)
	c.wr.Header().Set("Content-Type", ApplicationJSON)
}

func (c *Context) HandleError(err error) bool {
	if err != nil {
		ErrorF("HandleError: %v", err)
		c.err = err
		return true
	}
	return false
}

func (c *Context) Body(output any) error {
	ct := c.req.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		switch mediaType {
		case ApplicationJSON:
			dec := json.NewDecoder(c.req.Body)
			err := dec.Decode(output)
			return err
		}
	} else {
		return errors.New(Error_ContentTypeNotSet)
	}
	return nil
}
