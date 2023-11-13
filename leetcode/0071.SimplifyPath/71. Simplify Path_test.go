package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "test case 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "test case 2",
			path: "/../",
			want: "/",
		},
		{
			name: "test case 3",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "test case 4",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "test case 5",
			path: "/a/../../b/../c//.//",
			want: "/c",
		},
		{
			name: "test case 6",
			path: "/a//b////c/d//././/..",
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
