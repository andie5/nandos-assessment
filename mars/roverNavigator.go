package mars

import (
	"errors"
	"strconv"
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
