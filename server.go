package inutil

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

var server *Server_Model

func (ss *Start_Server) start() *Server_Model {
	server = &Server_Model{}

	if !strings.Contains(ss.Address, ":") {
		Debug("No port in address, using default 80")
		ss.Address = ss.Address + ":80"
	}

	ss.port = ":" + strings.Split(ss.Address, ":")[1]

	server.middleware_ch = &middleware_context_model{
		Contexts: map[*http.Request]*Context{},
	}

	server.Mux = http.NewServeMux()
	server.Mid = NewLogger(server.Mux)

	return server
}

func (s *Server_Model) Run() {
	internalLogF("Running server: %v", startModel.Server.port)
	// log.Fatal(http.ListenAndServe(startModel.Server.port, handler))

	log.Fatal(http.ListenAndServe(startModel.Server.port, s.Mid))
}

func (s *Server_Model) Get(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodGet, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Head(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodHead, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Post(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodPost, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Put(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodPut, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Patch(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodPatch, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Delete(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodDelete, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
}

func (s *Server_Model) Options(path string, h HandlerFunc) {
	internalLogF("Method: %v, Path: %v", MethodOptions, path)
	s.Mux.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
		h(s.Context(req))
	})
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
	internalLog("body")
	return nil
}
