package inutil

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareLog() HandlerFunc {
	return HandlerFunc(wrapperHandlerFromGin(gin.Logger()))
}

func MiddlewareRecovery() HandlerFunc {
	return HandlerFunc(func(c *Context) {

		defer func() {
			if r := recover(); r != nil {
				logInternalF("middlewareRecover: %v", r)
				c.JSON(ReturnInternalServerError("error: %v", r))
			}
		}()

		c.Next()

	})
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

func MiddlewareSafety() HandlerFunc {
	return HandlerFunc(func(c *Context) {
		c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Add("X-XSS-Protection", "1;mode=block")
		c.Writer.Header().Add("X-Frame-Options", "deny")
		c.Writer.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	})

}
