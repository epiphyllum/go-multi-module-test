package greet

import "testing"

func TestSayGreet(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

		{
			"1",
			args{"hewei"},
			"hello hewei",
		},
		{
			"2",
			args{"hewei"},
			"hello hewei",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayGreet(tt.args.str); got != tt.want {
				t.Errorf("SayGreet() = %v, want %v", got, tt.want)
			}
		})
	}
}
