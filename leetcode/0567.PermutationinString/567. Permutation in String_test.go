package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases
		{
			name: "Example 4",
			args: args{s1: "abcdxabcde", s2: "abcdeabcdx"},
			want: true,
		},
		{
			name: "Example 3",
			args: args{s1: "hello", s2: "ooolleoooleh"},
			want: false,
		},
		{
			name: "Example 1",
			args: args{s1: "ab", s2: "eidbaooo"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s1: "ab", s2: "eidboaoo"},
			want: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
