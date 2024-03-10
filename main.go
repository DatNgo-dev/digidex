package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/DatNgo-dev/digidex/api"
	"github.com/DatNgo-dev/digidex/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)

	fmt.Println("Server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
