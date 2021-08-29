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

func TestRoverStatus_processCommands(t *testing.T) {
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
			name: "Rover status is updated with the valid moves",
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
			name: "Rover status is returned as is if there are no commands to process",
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
			name: "If invalid move is in list of commands rover status is returned with this",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := RoverStatus{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			rover.processCommands(tt.args.commands, tt.args.rovers, tt.args.i, tt.args.grid)
			// Check rover status object after update
			if !reflect.DeepEqual(rover, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", rover, tt.want)
			}
		})
	}
}
