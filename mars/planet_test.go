package mars

import (
	"reflect"
	"testing"
)

func TestGetAxis(t *testing.T) {
	type args struct {
		dataInput []string
	}
	tests := []struct {
		name    string
		args    args
		want    PlanetAxis
		wantErr bool
	}{
		{
			name: "Returns an error if no data input is provided",
			args: args{
				dataInput: []string{},
			},
			want:    PlanetAxis{},
			wantErr: true,
		},
		{
			name: "Returns an error if the x or y coordinate is incorrectly formatted",
			args: args{
				dataInput: []string{
					"wrong format",
					"wrong format 2",
				},
			},
			want:    PlanetAxis{},
			wantErr: true,
		},
		{
			name: "Returns the planet axis object with the x and y coordinate if correct input is provided",
			args: args{
				dataInput: []string{
					"4 9",
					"1 2 N",
				},
			},
			want:    PlanetAxis{X: 4, Y: 9},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAxis(tt.args.dataInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAxis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAxis() = %v, want %v", got, tt.want)
			}
		})
	}
}
