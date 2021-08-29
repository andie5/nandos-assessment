package mars

import (
	"reflect"
	"testing"
)

func TestRoverStatus_rotateLeft(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction string
	}
	tests := []struct {
		name   string
		fields fields
		want   RoverStatus
	}{
		{
			name: "Updates rover status with the direction updated 90 degrees anticlockwise",
			fields: fields{
				X:         int(2),
				Y:         int(3),
				Direction: "N",
			},
			want: RoverStatus{X: int(2), Y: int(3), Direction: "W"},
		},
		{
			name:   "Default rover state is given with no direction if original rover object is empty",
			fields: fields{},
			want:   RoverStatus{X: int(0), Y: int(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			rover.rotateLeft()

			// Check rover status object after update
			if !reflect.DeepEqual(rover, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", rover, tt.want)
			}
		})
	}
}

func TestRoverStatus_rotateRight(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction string
	}
	tests := []struct {
		name   string
		fields fields
		want   RoverStatus
	}{
		{
			name: "Updates rover with the direction updated 90 degrees clockwise",
			fields: fields{
				X:         int(2),
				Y:         int(3),
				Direction: "E",
			},
			want: RoverStatus{X: int(2), Y: int(3), Direction: "S"},
		},
		{
			name:   "Default rover state is given with no direction if original rover object is empty",
			fields: fields{},
			want:   RoverStatus{X: int(0), Y: int(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			rover.rotateRight()

			// Check rover status object after update
			if !reflect.DeepEqual(rover, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", rover, tt.want)
			}
		})
	}
}

func TestRoverStatus_moveRover(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction string
	}
	tests := []struct {
		name   string
		fields fields
		want   RoverStatus
	}{
		{
			name: "Updates rover with one move west, in the x direction i.e -1 to current position",
			fields: fields{
				X:         int(1),
				Y:         int(3),
				Direction: "W",
			},
			want: RoverStatus{X: int(0), Y: int(3), Direction: "W"},
		},
		{
			name: "Updates rover with one move north, in the y direction i.e +1 to current position",
			fields: fields{
				X:         int(1),
				Y:         int(3),
				Direction: "N",
			},
			want: RoverStatus{X: int(1), Y: int(4), Direction: "N"},
		},
		{
			name:   "Default rover state is given with no direction if original rover object is empty",
			fields: fields{},
			want:   RoverStatus{X: int(0), Y: int(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			rover.moveRover()

			// Check rover status object after update
			if !reflect.DeepEqual(rover, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", rover, tt.want)
			}
		})
	}
}

func TestRoverStatus_validateMove(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction string
	}
	type args struct {
		grid PlanetAxis
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Returns true if rover coordinates is 0,0 which is within the planet grid axis",
			fields: fields{
				X:         int(0),
				Y:         int(0),
				Direction: "N",
			},
			args: args{
				grid: PlanetAxis{X: int(3), Y: int(1)},
			},
			want: true,
		},
		{
			name: "Returns true if rover coordinates is within the planet grid axis",
			fields: fields{
				X:         int(1),
				Y:         int(3),
				Direction: "N",
			},
			args: args{
				grid: PlanetAxis{X: int(3), Y: int(3)},
			},
			want: true,
		},
		{
			name: "Returns true if rover coordinates if one coordinate is within the planet grid axis and other coordinate is not provided i.e. 0",
			fields: fields{
				X:         int(1),
				Direction: "N",
			},
			args: args{
				grid: PlanetAxis{X: int(3), Y: int(1)},
			},
			want: true,
		},
		{
			name: "Returns false if rover coordinates is not within the planet grid axis",
			fields: fields{
				X:         int(1),
				Y:         int(3),
				Direction: "N",
			},
			args: args{
				grid: PlanetAxis{X: int(3), Y: int(1)},
			},
			want: false,
		},
		{
			name: "Returns false if planet grid axis is not supplied and the rover coordinates is not 0,0",
			fields: fields{
				X:         int(1),
				Y:         int(3),
				Direction: "N",
			},
			args: args{
				grid: PlanetAxis{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := &RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			if got := rover.validateMove(tt.args.grid); got != tt.want {
				t.Errorf("RoverStatus.validateMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
