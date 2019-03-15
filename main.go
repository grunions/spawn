package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Config struct {
	ListenAddr string
}

var config *Config

func init() {
	config = &Config{
		ListenAddr: os.Getenv("HTTP_LISTEN"),
	}
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	mux := chi.NewMux()
	mux.HandleFunc("/", debugHandler)
	http.ListenAndServe(config.ListenAddr, mux)
}
