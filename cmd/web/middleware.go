package main

import (
	"fmt"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("X-XSS-Protection", "1; mode=block")
			rw.Header().Set("X-Frame-Options", "deny")

			next.ServeHTTP(rw, r)
		},
	)
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

			next.ServeHTTP(rw, r)
		},
	)
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					rw.Header().Set("Connection", "close")

					app.serverError(rw, fmt.Errorf("%s", err))
				}
			}()

			next.ServeHTTP(rw, r)
		},
	)
}
