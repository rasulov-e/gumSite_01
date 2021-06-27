package main

import (
	"gumsite_01/server"
	"log"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.Start(":8080", "host=localhost user=postgres password=postgres dbname=local_server sslmode=disable"))
}
