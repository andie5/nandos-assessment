package main

import (
	"fmt"
	"log"
	"nandos/marsrover/v1/mars"
)

// FILENAME for data processing
const FILENAME = "./input/rovers.txt"

// Program starts here
func main() {

	// Read text input file
	fileOutput, err := mars.ReadFile(FILENAME)
	if err != nil {
		log.Fatalf("Failed to open file %v\n", err)
	}

	// Set planet grid axis
	planetGrid, err := mars.GetAxis(fileOutput)
	if err != nil {
		log.Fatalf("Error processing initial planet axis input: %v error: %v", fileOutput, err)
	}

	// Process rovers
	rovers := mars.ProcessRovers(fileOutput, planetGrid)

	// Print rover positions
	printRoverPositions(rovers)
}

// printRoverPositions prints the rover coordinates and direction
func printRoverPositions(rovers map[int]mars.RoverStatus) {
	for _, rover := range rovers {
		fmt.Printf("%v %v %v\n", rover.X, rover.Y, rover.Direction)
	}
}
