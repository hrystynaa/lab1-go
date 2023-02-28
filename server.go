package main

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	PATH = "/time"
	PORT = ":8795"
)

func main() {
	http.HandleFunc(PATH, realTimeHandler)
	http.ListenAndServe(PORT, nil)
}

func realTimeHandler(w http.ResponseWriter, r *http.Request) {
	realTime := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": realTime}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
