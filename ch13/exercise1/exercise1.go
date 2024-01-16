package main

import (
	"fmt"
	"net/http"
	"time"
)

// ----- 13.1 -----

func createServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))

		return
	})

	return mux
}

func main() {
	// ----- 13.1 -----
	fmt.Println("----- 13.1 -----")

	m := createServeMux()
	s := http.Server{
		Addr:         ":8080",
		Handler:      m,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
