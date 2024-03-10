package inutil

import "net/http"

func middleware_context_handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		server.middleware_ch.Contexts[req] = &Context{
			wr:   wr,
			req:  req,
			data: map[string]any{},
		}
		next.ServeHTTP(wr, req)
	})
}

func middleware_log_handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogF("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
