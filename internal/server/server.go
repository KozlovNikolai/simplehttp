package server

import (
	"fmt"
	"net/http"
)

type server struct {
	addr string
}

func New() *server {
	return &server{
		addr: "localhost:8080",
	}
}

func (s server) Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	fmt.Printf("Starting server at %s\n", s.addr)
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}
	return nil
}
