package middleware

import (
	"log"
	"net/http"
	"time"
)

type WrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func(w *WrappedWriter) WriteHeader(StatusCode int ){
	w.ResponseWriter.WriteHeader(StatusCode)
	w.statusCode =StatusCode
}

func Logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request){
		start := time.Now()

		wrapped := &WrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}