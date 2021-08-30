package mars

import (
	"reflect"
	"testing"
)

func Test_setStartingPosition(t *testing.T) {
	type args struct {
		rowContents []string
	}
	tests := []struct {
		name    string
		args    args
		want    RoverStatus
		wantErr bool
	}{
		{
			name: "Returns rover status object, given valid starting coordinates and string direction",
			args: args{
				rowContents: []string{
					"3",
					"4",
					"N",
				},
			},
			want: RoverStatus{
				X: int(3), Y: int(4), Direction: "N",
			},
			wantErr: false,
		},
		{
			name: "Returns an empty rover status object, and an error if input is empty",
			args: args{
				rowContents: []string{},
			},
			want:    RoverStatus{},
			wantErr: true,
		},
		{
			name: "Returns an empty rover status object, and an error if input is not the correct length of 3",
			args: args{
				rowContents: []string{"2", "3"},
			},
			want:    RoverStatus{},
			wantErr: true,
		},
		{
			name: "Returns an empty rover status object, and an error if input format is not valid",
			args: args{
				rowContents: []string{"A", "B", "W"},
			},
			want:    RoverStatus{},
			wantErr: true,
		},
		{
			name: "Returns an empty rover status object, and an error if input format is not valid",
			args: args{
				rowContents: []string{"A", "B", "W"},
			},
			want:    RoverStatus{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setStartingPosition(tt.args.rowContents)
			if (err != nil) != tt.wantErr {
				t.Errorf("setStartingPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setStartingPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoverStatus_runCommands(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction string
	}
	type args struct {
		commands string
		rovers   map[int]RoverStatus
		i        int
		grid     PlanetAxis
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RoverStatus
	}{
		{
			name: "Rover status is updated with the commands that result in valid move",
			fields: fields{
				X:         int(2),
				Y:         int(2),
				Direction: "N",
			},
			args: args{
				commands: "LLLLMR",
				rovers:   map[int]RoverStatus{},
				i:        int(4),
				grid: PlanetAxis{
					X: int(6), Y: int(6),
				},
			},
			want: RoverStatus{
				X:         int(2),
				Y:         int(3),
				Direction: "E",
			},
		},
		{
			name: "Rover status is returned if there are no commands to process",
			fields: fields{
				X:         int(2),
				Y:         int(2),
				Direction: "N",
			},
			args: args{
				commands: "",
				rovers:   map[int]RoverStatus{},
				i:        int(4),
				grid: PlanetAxis{
					X: int(6), Y: int(6),
				},
			},
			want: RoverStatus{
				X:         int(2),
				Y:         int(2),
				Direction: "N",
			},
		},
		{
			name: "If invalid command results in a invalid move then the rover status is returned immediately",
			fields: fields{
				X:         int(2),
				Y:         int(2),
				Direction: "N",
			},
			args: args{
				commands: "LLLLMMMMMR",
				rovers:   map[int]RoverStatus{},
				i:        int(4),
				grid: PlanetAxis{
					X: int(6), Y: int(6),
				},
			},
			want: RoverStatus{
				X:         int(2),
				Y:         int(7),
				Direction: "N",
			},
		},
		{
			name: "Only valid rover commands are processed, invalid commands are ignored",
			fields: fields{
				X:         int(2),
				Y:         int(2),
				Direction: "N",
			},
			args: args{
				commands: "LLKAM",
				rovers:   map[int]RoverStatus{},
				i:        int(4),
				grid: PlanetAxis{
					X: int(6), Y: int(6),
				},
			},
			want: RoverStatus{
				X:         int(2),
				Y:         int(1),
				Direction: "S",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			rover.runCommands(tt.args.commands, tt.args.rovers, tt.args.i, tt.args.grid)
			// Check rover status object after update
			if !reflect.DeepEqual(rover, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", rover, tt.want)
			}
		})
	}
}

func TestProcessRovers(t *testing.T) {
	type args struct {
		dataInput []string
		grid      PlanetAxis
	}
	tests := []struct {
		name string
		args args
		want map[int]RoverStatus
	}{
		{
			name: "Returns list of rovers and their final positions",
			args: args{
				dataInput: []string{
					"5 5",
					"1 2 N",
					"LMLMLMLMM",
					"3 3 E",
					"MMRMMRMRRM",
				},
				grid: PlanetAxis{
					X: int(5), Y: int(5),
				},
			},
			want: map[int]RoverStatus{
				1: {
					X:         int(1),
					Y:         int(3),
					Direction: "N",
				},
				3: {
					X:         int(5),
					Y:         int(1),
					Direction: "E",
				},
			},
		},
		{
			name: "Returns empty rovers if no data input is provided",
			args: args{
				dataInput: []string{},
				grid: PlanetAxis{
					X: int(5), Y: int(5),
				},
			},
			want: map[int]RoverStatus{},
		},
		{
			name: "Returns empty list of rovers if planet grid axis is not initialised",
			args: args{
				dataInput: []string{
					"5 5",
					"1 2 N",
					"LMLMLMLMM",
					"3 3 E",
					"MMRMMRMRRM",
				},
				grid: PlanetAxis{},
			},
			want: map[int]RoverStatus{},
		},
		{
			name: "Returns empty rover status for any rovers where the starting position is invalid",
			args: args{
				dataInput: []string{
					"5 5",
					"1 2",
					"LMLMLMLMM",
					"3 3 E",
					"MMRMMRMRRM",
				},
				grid: PlanetAxis{
					X: int(5), Y: int(5),
				},
			},
			want: map[int]RoverStatus{
				1: {},
				3: {
					X:         int(5),
					Y:         int(1),
					Direction: "E",
				},
			},
		},
		{
			name: "Returns list of rovers and their final positions. Rovers can start and end in the same position",
			args: args{
				dataInput: []string{
					"5 5",
					"1 2 N",
					"LMLMLMLMM",
					"1 2 N",
					"MLLLL",
					"1 4 S",
					"MMML",
				},
				grid: PlanetAxis{
					X: int(5), Y: int(5),
				},
			},
			want: map[int]RoverStatus{
				1: {
					X:         int(1),
					Y:         int(3),
					Direction: "N",
				},
				3: {
					X:         int(1),
					Y:         int(3),
					Direction: "N",
				},
				5: {
					X:         int(1),
					Y:         int(1),
					Direction: "E",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessRovers(tt.args.dataInput, tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessRovers() = %v, want %v", got, tt.want)
			}
		})
	}
}
