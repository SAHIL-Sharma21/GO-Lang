package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SAHIL-Sharma21/mongoAPI/routers"
)

func main() {
	fmt.Println("MongoDB API")

	r := routers.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is listening at port 4000.....")
}
