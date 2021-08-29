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
