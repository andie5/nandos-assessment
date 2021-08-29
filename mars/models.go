package mars

// PlanetAxis is the grid axis for the overall planet
type PlanetAxis struct {
	X int
	Y int
}

// RoverStatus stores the status of the rover, the x, y coordinates and the direction the rover is in
type RoverStatus struct {
	X         int
	Y         int
	Direction string
}
