# Nandos assessment - MARS ROVER exercise

27/08/2021 - Andrea Asamoah

## Summary

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.
This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.

A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points.

The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North.

In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.

'M' means move forward one grid point, and maintain the same heading.
Assume that the square directly North from (x, y) is (x, y+1).

## Input:

The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be 0,0.
The rest of the input is information pertaining to the rovers that have been deployed. Each rover has two lines of input.

The first line gives the rover's position, and the second line is a series of instructions telling the rover how to explore the plateau.

The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's orientation.

Each rover will be finished sequentially, which means that the second rover won't start to move until the first one has finished moving.

# Output:

The output for each rover should be its final co-ordinates and heading.

## Test Input:

5 5  
1 2 N  
LMLMLMLMM  
3 3 E  
MMRMMRMRRM

## Expected Output:

1 3 N  
5 1 E

# Getting started

- Go 1.16.5 has been used for development of this application
- The input file is contained in the input folder. This input file is needed to run the program: `rovers.txt`
- Update this file if you want to add more test examples

## Run the program

To run the solution, with Go:

- `go run .` or `go build .` followed by `./mars`

To run the solution, using Docker:

- `docker build --tag mars .`
- Then `docker run --name marsrover -it mars:latest` to see output

# Assumptions

- If any input directions are not "L", "M" or "R", then that direction is just ignored for a given rover.
- Rovers can start and finish in the same location.
- The planet grid size coordinates and the rover coordinates will fit into the default 32 bits or 64 bits dependent on the platform. Specific int base has not been considered as part of this implementation
- The input structure will always have the grid size as the first row, then alternate rows for the starting position of the rover followed by the commands for that rover.
- All commands for rover moves provided are uppercase.
- All starting positions for rovers are uppercase.

# Future considerations

- The input file could be another process or a database, where connection details would become part of a .env file.
- Each rover could have a unique id created as part of a function, to be retrieved later.
- The maps type that stores the orientation details could be changed to actual functions for more complex rotations.
