package main

import (
	"fmt"
	"github.com/ddvalim/go-pdf-builder/router"
	"log"
	"net/http"
)

const port = 8080

func main() {
	fmt.Println(fmt.Sprintf("server running at %d", port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router.New())
	if err != nil {
		log.Fatal(err)
	}
}
