package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// ----- 13.2 -----

func createServeMuxForLoggingIP(logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		ip, _, _ := strings.Cut(r.RemoteAddr, ":")
		logger.Info("incoming IP", "ip", ip)

		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))

		return
	})

	return mux
}

func main() {
	// ----- 13.2 -----
	fmt.Println("----- 13.2 -----")

	options := &slog.HandlerOptions{}
	handler := slog.NewJSONHandler(os.Stderr, options)
	log := slog.New(handler)
	mux := createServeMuxForLoggingIP(log)
	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
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
