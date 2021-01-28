package main

import (
	"fmt"
	"urlShortener/api"
)

func main() {
	fmt.Println("Server started")
	api.HandleRequests()
}

//TODO: add more logs (one link/key)
