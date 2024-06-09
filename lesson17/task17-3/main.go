package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("log")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	l := log.New(f, "", 0)
	logHandler := logMiddleware(l)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello(l))

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux), l),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(l *log.Logger) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		msg := "Hello, Go!"
		l.Println("resp:", msg)
		_, err := fmt.Fprintf(res, msg)
		if err != nil {
			l.Println(err)
		}
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
