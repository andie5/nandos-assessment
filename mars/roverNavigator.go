package mars

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// setStartingPosition creates a rover status object withe starting coorindates and orientation direction and assigns it the master rover struct
func setStartingPosition(rowContents []string) (RoverStatus, error) {
	// Check the X, Y and orientation is provided
	if len(rowContents) != 3 {
		return RoverStatus{}, errors.New("rover starting position is invalid - missing entries")
	}
	xValue, err := strconv.Atoi(rowContents[0])
	if err != nil {
		return RoverStatus{}, err
	}
	yValue, err := strconv.Atoi(rowContents[1])
	if err != nil {
		return RoverStatus{}, err
	}

	// Create rover status object
	return RoverStatus{
		X:         xValue,
		Y:         yValue,
		Direction: rowContents[2],
	}, nil
}

// processCommands updates the rover position based on a list of string commands. Following a command, the rover position is checked to see if it is valid.
// If the rover move position is invalid at any point, the rover status is returned as is and no further moves are processed
func (rover *RoverStatus) processCommands(commands string, rovers map[int]RoverStatus, i int, grid PlanetAxis) {
	for _, command := range commands {
		validMove := true // assume move is valid unless set otherwise
		switch string(command) {
		case "M":
			rover.moveRover()
			validMove = rover.validateMove(grid) // validate move is within grid axis
		case "R":
			rover.rotateRight()
		case "L":
			rover.rotateLeft()
		default:
			// Unknown command, no update
		}
		rovers[i-1] = *rover
		fmt.Println("command: ", string(command), " rover:", rover.X, " ", rover.Y, " ", rover.Direction)
		if !validMove {
			log.Printf("Invalid rover move. Rover outside of planet grid axis: %v rover: %v", grid, rover)
			break // Don't process any more moves for this rover
		}
	}
}

// ProcessRovers loops through the data input and updates the rover postion based on a starting position and list of commands
func ProcessRovers(dataInput []string, grid PlanetAxis) map[int]RoverStatus {

	// Initialise map of rovers
	rovers := map[int]RoverStatus{}
	// If planet axis coordinates have not been set or are both zero, return false
	if grid.X == 0 && grid.Y == 0 {
		return rovers
	}

	for i := 1; i < len(dataInput); i++ {
		row := dataInput[i]
		rowContents := strings.Split(row, " ")
		// Every other row of the data is the rover directions
		if i%int(2) == 0 {
			// Retrieve rover and process rover commands, one by one
			rover := rovers[i-1]
			rover.processCommands(row, rovers, i, grid)
		} else {
			var err error
			rovers[i], err = setStartingPosition(rowContents)
			// Check if starting position is in the grid, if not, skip to next rover
			if err != nil {
				log.Printf("Cannot set starting position for rover. Conversion error: %v\n", err)
				i += 1 // Skip the commands for the rover
				continue
			}
		}
	}
	return rovers
}
