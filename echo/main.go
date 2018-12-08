package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	serverPort = "8080"
)

type Result struct {
	Meta string `json:"meta"`
	Msg  string `json:"msg"`
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil); err != nil {
		log.Fatal("failed to serve http server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Meta: "v1",
		Msg:  r.URL.Query().Get("msg"),
	}

	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
