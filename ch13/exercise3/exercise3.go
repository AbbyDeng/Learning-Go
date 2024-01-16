package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// ----- 13.3 -----

func buildText(now time.Time) string {
	return now.Format(time.RFC3339)
}

func buildJSON(now time.Time) string {
	timeOut := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}

func createServeMuxForReturningJSON(logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		ip, _, _ := strings.Cut(r.RemoteAddr, ":")
		logger.Info("incoming IP", "ip", ip)

		now := time.Now()
		var output string
		if r.Header.Get("Accept") == "application/json" {
			output = buildJSON(now)
		} else {
			output = buildJSON(now)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))

		return
	})

	return mux
}

func main() {
	// ----- 13.3 -----
	fmt.Println("----- 13.3 -----")

	options := &slog.HandlerOptions{}
	handler := slog.NewJSONHandler(os.Stderr, options)
	log := slog.New(handler)
	mux := createServeMuxForReturningJSON(log)
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
