package main

import (
	"log"
)

func doStuff() string {
	return "I do stuff!"
}

func mainPr() {
	log.Printf("starting service sonar ...")
	log.Printf(doStuff())
}
