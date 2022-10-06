package controller

import (
	"log"
	"net/http"
	"time"
)

func AllowOnlyPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteJsonResponse(w, &Response{
				Status:  400,
				Message: "BAD_REQUEST",
				Error:   "method" + r.Method + "not allowed",
			})
			return
		}
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			WriteJsonResponse(w, &Response{
				Status:  400,
				Message: "BAD_REQUEST",
				Error:   "need login!",
			})
			return
		}
		if username != "admin" || password != "admin" {
			WriteJsonResponse(w, &Response{
				Status:  400,
				Message: "BAD_REQUEST",
				Error:   "Username or password not matched",
			})
			return
		}

		next(w, r)
	}
}

func Trace(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next(w, r)

		end := time.Since(now)

		log.Println("response time", end.Milliseconds())
	}
}
