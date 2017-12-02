package day1

import (
	"testing"
)

func Test_reverseCaptcha(t *testing.T) {
	type args struct {
		inputNumbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1122",
			args: args{
				inputNumbers: []int{
					1,
					1,
					2,
					2,
				},
			},
			want: 3,
		},
		{
			name: "1111",
			args: args{
				inputNumbers: []int{
					1,
					1,
					1,
					1,
				},
			},
			want: 4,
		},
		{
			name: "1234",
			args: args{
				inputNumbers: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: 0,
		},
		{
			name: "91212129",
			args: args{
				inputNumbers: []int{
					9,
					1,
					2,
					1,
					2,
					1,
					2,
					9,
				},
			},
			want: 9,
		},
		{
			name: "349977448929146532968278716",
			args: args{
				inputNumbers: []int{
					3, 4, 9, 9, 7, 7, 4, 4, 8, 9, 2, 9, 1, 4, 6, 5, 3, 2, 9, 6, 8, 2, 7, 8, 7, 1, 6,
				},
			},
			want: 9 + 7 + 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseCaptcha(tt.args.inputNumbers); got != tt.want {
				t.Errorf("reverseCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverseCaptchaPartTwo(t *testing.T) {
	type args struct {
		inputNumbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1122",
			args: args{
				inputNumbers: []int{
					1,
					1,
					2,
					2,
				},
			},
			want: 0,
		},
		{
			name: "1111",
			args: args{
				inputNumbers: []int{
					1,
					1,
					1,
					1,
				},
			},
			want: 4,
		},
		{
			name: "1234",
			args: args{
				inputNumbers: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: 0,
		},
		{
			name: "91212129",
			args: args{
				inputNumbers: []int{
					9,
					1,
					2,
					1,
					2,
					1,
					2,
					9,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseCaptchaPartTwo(tt.args.inputNumbers); got != tt.want {
				t.Errorf("reverseCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
