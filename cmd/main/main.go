package main

import (
	"log"
	"myMongoTest/pkg/config/router"
	"net/http"
)

func main() {
	// creating router
	r := router.Router()
	// making router work, till user stops it
	log.Fatal(http.ListenAndServe(":4000", r))

}
