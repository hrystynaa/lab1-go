package main

import (
	"encoding/json"
	"log"
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
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func realTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error: you can use only GET method", http.StatusNotFound)
		return
	}
	realTime := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": realTime}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
