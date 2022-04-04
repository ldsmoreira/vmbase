package main

import (
	"log"
	"modalchemy-virt-plataform/api"
	"net/http"
)

func main() {

	http.HandleFunc("/", api.HandleAPIResquest)
	log.Fatal(http.ListenAndServe(":20300", nil))

}
