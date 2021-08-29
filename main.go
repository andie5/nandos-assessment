package main

import (
	"fmt"
	"log"
	"nandos/marsrover/v1/mars"
)

const (
	FILENAME = "./input/rovers.txt"
)

// Program starts here
func main() {

	// Read text input file
	fileOutput, err := mars.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("file did not process %v %v\n", FILENAME, err)
	}

	// Set planet grid axis
	planetGrid, err := mars.GetAxis(fileOutput)
	if err != nil {
		log.Fatalf("Error processing initial planet axis input: %v error: %v", fileOutput, err)
	}

	//temp initialise rover object
	rovers := map[int]mars.RoverStatus{}
	printRoverPositions(rovers)

	log.Println("filename: ", fileOutput)
	log.Println("planetGrid: ", planetGrid)

}

// printRoverPositions prints the rover coordinates and direction
func printRoverPositions(rovers map[int]mars.RoverStatus) {
	for _, rover := range rovers {
		fmt.Printf("%v %v %v\n", rover.X, rover.Y, rover.Direction)
	}
}
