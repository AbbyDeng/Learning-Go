package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// ----- 14.1 -----

func Timeout(t int) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Duration(t)*time.Microsecond)
			defer cancel()

			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func doSomething(ctx context.Context) (string, error) {
	wait := rand.Intn(200)
	select {
	case <-time.After(time.Duration(wait) * time.Microsecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too Slow!", ctx.Err()
	}
}

func sleepy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msg, err := doSomething(ctx)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(msg))
}

func main() {
	// ----- 14.1 -----
	fmt.Println("----- 14.1 -----")

	middleware := Timeout(100)
	server := http.Server{
		Handler: middleware(http.HandlerFunc(sleepy)),
		Addr:    ":8080",
	}

	server.ListenAndServe()
}
