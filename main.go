package main

import (
	"log"
	"nandos/marsrover/v1/mars"
)

const (
	FILENAME = "./input/rovers.txt"
)

// Program starts here
func main() {

	// Read text input file
	file, err := mars.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("file did not process %v %v\n", FILENAME, err)
	}

	log.Println("filename: ", file)

}
