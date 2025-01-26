package model

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
	gc *gin.Context
}
