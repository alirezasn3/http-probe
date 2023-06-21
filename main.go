package main

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	if err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if !isValidIP(string(reqBody)) {
			w.WriteHeader(400)
			return
		}
		_, err = client.Get("http://" + string(reqBody) + ":80")
		if err != nil && !strings.Contains(strings.ToLower(err.Error()), "refused") {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})); err != nil {
		panic(err)
	}
}
