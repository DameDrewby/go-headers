package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: No command line arguments provided")
		fmt.Println("Usage: query_webserver <hostname> ")
		os.Exit(1)
	}

	res, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The following headers were retrieved")

	for name, value := range res.Header {
		fmt.Printf("%s  %s\n", name, value[0])
	}
}
