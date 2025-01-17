package __17

import "testing"

func Test_minimumSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 1, 8},
				k:    10,
			},
			wantRet: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 32, 21},
				k:    55,
			},
			wantRet: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := minimumSubarrayLength(tt.args.nums, tt.args.k); gotRet != tt.wantRet {
				t.Errorf("minimumSubarrayLength() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
