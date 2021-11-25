package main

import (
	"github.com/kaanozkan/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
