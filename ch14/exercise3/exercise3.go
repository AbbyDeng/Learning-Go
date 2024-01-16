package main

import (
	"errors/ch14/exercise3/log"
	"fmt"
	"net/http"
)

// ----- 14.3 -----

func message(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Log(ctx, log.Debug, "This is a debug message")
	log.Log(ctx, log.Info, "This is an info message")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}

func main() {
	// ----- 14.3 -----
	fmt.Println("----- 14.3 -----")

	server := http.Server{
		Handler: log.Middleware(http.HandlerFunc(message)),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
