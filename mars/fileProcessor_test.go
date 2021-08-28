package mars

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Returns an error if no filename is present",
			args: args{
				filename: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Returns an error if file cannot be found",
			args: args{
				filename: "unknown.txt",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Returns an error if file cannot be found",
			args: args{
				filename: "../input/test.txt",
			},
			want: []string{
				"abc",
				"def",
				"123",
				"456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
