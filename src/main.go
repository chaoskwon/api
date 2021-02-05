package main

import (
	// "fmt"
	// "github.com/Sirupsen/logrus"

	"route"	
)

// func init() {
// 	db = lib.NewConnection()
// 	// logrus.SetLevel(logrus.DebugLevel)
// 	// logrus.SetFormatter(&logrus.JSONFormatter{})
// }

func main() {
	// router := route.Init()
	router := route.Router()
	// Start server
	router.Start(":1323")
}
