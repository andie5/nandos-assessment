package mars

import (
	"errors"
	"strconv"
	"strings"
)

// GetAxis gets the axis details for the planet
func GetAxis(dataInput []string) (PlanetAxis, error) {

	planet := PlanetAxis{}

	if len(dataInput) != 0 { // Check if there is planet coordinates provided
		firstRow := dataInput[0]
		rowContents := strings.Split(firstRow, " ")

		if len(rowContents) != 2 {
			return planet, errors.New("planet axis not provided in correct format")
		}
		xValue, err := strconv.Atoi(rowContents[0])
		if err != nil {
			return planet, err
		}
		yValue, err := strconv.Atoi(rowContents[1])
		if err != nil {
			return planet, err
		}
		return PlanetAxis{xValue, yValue}, nil
	}
	// Default if file is empty
	return planet, errors.New("no content to process")
}
