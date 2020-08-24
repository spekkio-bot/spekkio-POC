package app

import (
	"log"
	"net"
	"net/http"
)

/*
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
*/

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = "-"
		}

		handler.ServeHTTP(w, r)
		/*
		lrw := NewLoggingResponseWriter(w)
		handler.ServeHTTP(lrw, r)

		status := lrw.statusCode
		*/
		log.Printf("%s \"%s %s %s\"", ip, r.Method, r.RequestURI, r.Proto)
	})
}
