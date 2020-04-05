// log.go

package main

import (
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	logInfo("Request: %s %s", r.Method, r.RequestURI)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logRequest(r)
        next.ServeHTTP(w, r)
    })
}

func logInfo(format string, v ...interface{}) {
	logFull("INFO", format, v...)
}

func logError(format string, v ...interface{}) {
	logFull("ERROR", format, v...)
}

func logFull(level string, format string, v ...interface{}) {
	format = level + ": " + format
	log.Printf(format, v...)
}