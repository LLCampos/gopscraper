package main

import (
    "io"
    "net/http"

	"gopscraper"
)

func contestsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    io.WriteString(w, gopscraper.GetContests())
}

func main() {
    http.HandleFunc("/passatempos", contestsHandler)
    http.ListenAndServe(":5001", nil)
}
