package main

import (
	"go-banking/app"
	"go-banking/log"
)

func main() {
	
	// log.Println("Starting our application...")
	log.Info("Startting the application...")
	app.Start()
}