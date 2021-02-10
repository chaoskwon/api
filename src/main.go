package main

import (
	// "fmt"

	"route"	
)

func main() {
	router := route.Router()
	// Start server
	router.Logger.Fatal(router.Start(":1323"))
}
