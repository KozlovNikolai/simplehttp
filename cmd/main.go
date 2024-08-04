package main

import (
	"log"

	"github.com/KozlovNikolai/simplehttp/internal/server"
)

func main() {
	s := server.New()
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
