package inutil

import "net/http"

func middleware_context_handler(next http.Handler) http.Handler {
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

		wr.Header().Add("Access-Control-Allow-Origin", "*")
		// wr.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(wr, req)
	})
}
