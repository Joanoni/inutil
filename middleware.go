package inutil

import "net/http"

func middleware_log_handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		LogF("%s - %s (%s)", req.Method, req.URL.Path, req.RemoteAddr)
		req.Context()
		next.ServeHTTP(wr, req)
	})
}

func middleware_cors_handler(c HandlerFunc) HandlerFunc {
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
