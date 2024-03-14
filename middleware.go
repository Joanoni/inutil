package inutil

import "net/http"

func middleware_context_handler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		c := &Context{
			wr:   wr,
			req:  req,
			data: map[string]any{},
		}
		server.middleware_ch.Contexts[req] = c
		next.ServeHTTP(wr, req)
		if c.err != nil {
			c.JSON(Return[any]{
				Message: c.err.Error(),
				Data:    nil,
				Success: false,
				Status:  StatusBadRequest,
			})
		}
	})
}

func middleware_log_handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		LogF("%s - %s (%s)", req.Method, req.URL.Path, req.RemoteAddr)

		next.ServeHTTP(wr, req)
	})
}

func middleware_cors_handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		internalLog("entrou")
		wr.Header().Add("Access-Control-Allow-Origin", "*")
		wr.Header().Add("Access-Control-Allow-Credentials", "true")
		wr.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		wr.Header().Add("Access-Control-Allow-Methods", "POST,OPTIONS,GET,PUT,DELETE,PATCH")
		wr.Header().Add("X-Content-Type-Options", "nosniff")
		wr.Header().Add("X-XSS-Protection", "1;mode=block")
		wr.Header().Add("X-Frame-Options", "deny")
		wr.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		if req.Method == "OPTIONS" {
			wr.WriteHeader(StatusNoContent)
			return
		}
		internalLog("passou")
		next.ServeHTTP(wr, req)
	})
}
