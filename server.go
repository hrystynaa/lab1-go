package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", realTimeHandler)
	http.ListenAndServe(":8795", nil)
}

func realTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"time": time.Now().Format(time.RFC3339)})
}
