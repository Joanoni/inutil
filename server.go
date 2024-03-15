package inutil

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

var server *Server_Model

func (ss *Start_Server) start() *Server_Model {
	server = &Server_Model{}

	if ss.Port == "" {
		Debug("No port in address, using default :80")
		server.port = ":80"
	} else {
		server.port = ss.Port
	}

	server.engine = gin.Default()

	server.engine.Use(gin.Logger())

	return server
}

func wrapperHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&Context{
			ginContext: c,
		})
	}
}

func (s *Server_Model) Run() error {
	internalLogF("Running server: %v", s.port)

	return s.engine.Run(s.port)
}

func (s *Server_Model) Get(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Head(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Post(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Put(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Patch(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Delete(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (s *Server_Model) Options(path string, h HandlerFunc) {
	s.engine.GET(path, wrapperHandler(h))
}

func (c *Context) JSON(payload Return[any]) {
	c.ginContext.JSON(payload.Status, payload)
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
	ct := c.ginContext.Request.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		switch mediaType {
		case ApplicationJSON:
			dec := json.NewDecoder(c.ginContext.Request.Body)
			err := dec.Decode(output)
			return err
		}
	} else {
		return errors.New(Error_ContentTypeNotSet)
	}
	internalLog("body")
	return nil
}
