package inutil

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *Context)

func wrapperHandlersToGin(hfs ...HandlerFunc) []gin.HandlerFunc {
	ghfs := []gin.HandlerFunc{}
	for _, hf := range hfs {
		ghfs = append(ghfs, func(c *gin.Context) {
			hf(convertContextFromGin(c))
		})
	}
	return ghfs
}

func wrapperHandlerToGin(hf HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		hf(convertContextFromGin(c))
	}
}

func wrapperHandlersFromGin(ghfs ...gin.HandlerFunc) []HandlerFunc {
	hfs := []HandlerFunc{}
	for _, ghf := range ghfs {
		hfs = append(hfs, func(c *Context) {
			ghf(convertContextToGin(c))
		})
	}
	return hfs
}

func wrapperHandlerFromGin(ghf gin.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		ghf(convertContextToGin(c))
	}
}
