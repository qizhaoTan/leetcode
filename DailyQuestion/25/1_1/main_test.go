package main

import "testing"

func Test_convertDateToBinary(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"1", "2080-02-29", "100000100000-10-11101"},
		{"2", "1900-01-01", "11101101100-1-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDateToBinary(tt.args); got != tt.want {
				t.Errorf("convertDateToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
