package server

import (
	"github.com/Joanoni/inutil"
	"github.com/gin-gonic/gin"
)

var server *Server

func (ssi *StartServerInput) start() *Server {
	server = &Server{}

	if ssi.Port == "" {
		Print("No port in address, using default :80")
		server.port = ":80"
	} else {
		server.port = ssi.Port
	}

	server.engine = gin.New()

	return server
}

func (s *Server) Run() error {
	logInternalF("Running server: %v", s.port)

	if inutil.WebSocketManager != nil {
		s.Get(inutil.WebSocketManager.path, WebsocketHandler())
	}

	return s.engine.Run(s.port)
}

func (s *Server) Use(handlers ...HandlerFunc) {
	s.engine.Use(wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Get(path string, handlers ...HandlerFunc) {
	s.engine.GET(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Head(path string, handlers ...HandlerFunc) {
	s.engine.HEAD(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Post(path string, handlers ...HandlerFunc) {
	s.engine.POST(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Put(path string, handlers ...HandlerFunc) {
	s.engine.PUT(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Patch(path string, handlers ...HandlerFunc) {
	s.engine.PATCH(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Delete(path string, handlers ...HandlerFunc) {
	s.engine.DELETE(path, wrapperHandlersToGin(handlers...)...)
}

func (s *Server) Options(path string, handlers ...HandlerFunc) {
	s.engine.OPTIONS(path, wrapperHandlersToGin(handlers...)...)
}
