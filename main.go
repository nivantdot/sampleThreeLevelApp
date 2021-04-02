package main

import (
	"log"
	"sampleThreeLevelApp/controller"
)

func myController() {
	controller.Controller()
}

func main() {
	log.Println("running on localhost:12345")
	myController()
}
