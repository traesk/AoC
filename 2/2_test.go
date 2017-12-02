package day2

import (
	"reflect"
	"testing"
)

func Test_makeSpreadSheet(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				file: "./test",
			},
			want: [][]int{
				{
					5, 1, 9, 5,
				},
				{
					7, 5, 3,
				},
				{
					2, 4, 6, 8,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, err := makeSpreadSheet(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSpreadSheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSpreadSheet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve1(tt.args.path); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getChecksum(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "testcase 1",
			args: args{
				path: "./test1",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "testcase 2",
			args: args{
				path: "./test2",
			},
			want:    160,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getChecksum(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_Solve2(t *testing.T) {
	a := Solve2("test3")
	if a != 9 {
		t.Errorf("Solve2_Test: want %v, got %v", 9, a)
	}
}
