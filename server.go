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

	return server
}

func wrapperHandler(hfs ...HandlerFunc) []gin.HandlerFunc {
	whfs := []gin.HandlerFunc{}
	for _, hf := range hfs {
		whfs = append(whfs, func(c *gin.Context) {
			hf(&Context{
				ginContext: c,
			})
		})
	}
	return whfs
}

func (s *Server_Model) Run() error {
	internalLogF("Running server: %v", s.port)

	return s.engine.Run(s.port)
}

func (s *Server_Model) Use(handlers ...HandlerFunc) {
	s.engine.Use(wrapperHandler(handlers...)...)
}

func (s *Server_Model) Get(path string, handlers ...HandlerFunc) {
	s.engine.GET(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Head(path string, handlers ...HandlerFunc) {
	s.engine.HEAD(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Post(path string, handlers ...HandlerFunc) {
	s.engine.POST(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Put(path string, handlers ...HandlerFunc) {
	s.engine.PUT(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Patch(path string, handlers ...HandlerFunc) {
	s.engine.PATCH(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Delete(path string, handlers ...HandlerFunc) {
	s.engine.DELETE(path, wrapperHandler(handlers...)...)
}

func (s *Server_Model) Options(path string, handlers ...HandlerFunc) {
	s.engine.OPTIONS(path, wrapperHandler(handlers...)...)
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
