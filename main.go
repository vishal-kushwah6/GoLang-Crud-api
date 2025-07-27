package main

import (
	"fmt"
	"log"
	"net/http"
	"netflix-clone/router"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	fmt.Println("server starting...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server started at port = 4000")

}
