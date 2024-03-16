package inutil

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareLog() HandlerFunc {
	return HandlerFunc(wrapperHandlerFromGin(gin.Logger()))
}

func MiddlewareCors() HandlerFunc {
	return HandlerFunc(func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

}

// wr.Header().Add("X-Content-Type-Options", "nosniff")
// wr.Header().Add("X-XSS-Protection", "1;mode=block")
// wr.Header().Add("X-Frame-Options", "deny")
// wr.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
