package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
