package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func ip(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://ip.mrfriday.com", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
  defer res.Body.Close()
  io.Copy(w, res.Body)
}

func health(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  w.Write([]byte("ok\n"))
}

func main() {
	http.HandleFunc("/ip", ip)
  http.HandleFunc("/health", health)
	log.Println("about to listen on port 8989")
	http.ListenAndServe(":8989", nil)
}
