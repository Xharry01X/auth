package main

import (
	"github.com/go-fuego/fuego"
)

func main() {
	// Create a new server instance with a specific address
	s := fuego.NewServer(fuego.WithAddr(":8080"))

	// Define a GET route
	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	// Run the server
	s.Run()
}
