package main

import (
	"log"

	"github.com/Narutchai01/ProjectC-BE/api"
	"github.com/Narutchai01/ProjectC-BE/db"
	_ "github.com/lib/pq"
)

func main() {
	store, err := db.NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
