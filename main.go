package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)
	AddHandlers(router)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", router)
}
