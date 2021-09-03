package main

import (
	"log"
	"net/http"

	"restapidemo/routers"
)

func main() {
	routers.Router.HandleFunc("/", welcomeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", routers.Router))
}

func welcomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Welcome to API"))
}
