package mars

// DirectionToMove maps the direction to the positive or negative move of 1 or -1
var DirectionToMove = map[string]int{
	"N": 1,
	"E": 1,
	"S": -1,
	"W": -1,
}

// rotateLeft updates the rover position 90 degrees in a anti-clockwise position N, E, S or W
func (rover *RoverStatus) rotateLeft() {
	anticlockwiseOrientation := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	rover.Direction = anticlockwiseOrientation[rover.Direction]
}

// rotateRight updates the rover position 90 degrees in a clockwise position N, E, S or W
func (rover *RoverStatus) rotateRight() {
	clockwiseOrientation := map[string]string{
		"N": "E",
		"W": "N",
		"S": "W",
		"E": "S",
	}
	rover.Direction = clockwiseOrientation[rover.Direction]
}

// moveRover updates the rover x or y position based on the direction
func (rover *RoverStatus) moveRover() {
	move := DirectionToMove[rover.Direction]
	if rover.Direction == "N" || rover.Direction == "S" {
		rover.Y += move
	} else if rover.Direction == "E" || rover.Direction == "W" {
		rover.X += move
	}
}

// validateMove checks that the rover's latest move position is within the grid axis coordinates and returns true if it is and false otherwise
func (rover *RoverStatus) validateMove(grid PlanetAxis) bool {
	// If planet axis coordinates have not been set or are both zero, return false
	if grid.X == 0 && grid.Y == 0 {
		return false
	}
	if (rover.X <= grid.X && rover.X >= 0) && (rover.Y <= grid.Y && rover.Y >= 0) {
		return true
	}
	return false
}
