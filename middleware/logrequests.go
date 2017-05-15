package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//shared function called from each handler
func logReq(r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

//single function closure approach
func logReqs(hfn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		start := time.Now()
		hfn(w, r)
		fmt.Printf("%v\n", time.Since(start))
	}
}

//Adapter style middleware
func logRequests(logger *log.Logger) Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s", r.Method, r.URL.Path)
			start := time.Now()
			handler.ServeHTTP(w, r)
			logger.Printf("%v\n", time.Since(start))
		})
	}
}
