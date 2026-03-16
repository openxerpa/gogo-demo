package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"service":     "gogo-demo",
			"status":      "ok",
			"version":     "1.0.0",
			"timestamp":   time.Now().Format(time.RFC3339),
			"environment": env,
		})
	})
	http.ListenAndServe(":"+port, nil)
}
