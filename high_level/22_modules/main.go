package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("modules in go lang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	//server running
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey Mod users")
}

// r is requst and its a pointer and w is response writer
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go lang server</h1>"))
}
