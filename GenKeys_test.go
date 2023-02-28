package go_utils

import "testing"

func TestGenKeys(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test1",
			args{"5BE0D6D4-E052-40B4-937C-76914C2909B3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenKeys(tt.args.input)
		})
	}
}
