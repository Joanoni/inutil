package model

import "github.com/gin-gonic/gin"

type Server struct {
	port   string
	engine *gin.Engine
}

type StartServerInput struct {
	Port string
}
