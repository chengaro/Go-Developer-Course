package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	l := log.New(os.Stdout, "", 0)

	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux), l),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	_, err := fmt.Fprintf(res, msg)
	if err != nil {
		log.Println(err)
	}
}

func authHandler(h http.Handler, l *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			l.Println("unauthorized request to:", r.URL)
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
