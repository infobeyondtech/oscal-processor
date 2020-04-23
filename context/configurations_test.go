package Context

import (
	"os/user"
	"testing"
)

func TestExpandPath(t *testing.T) {
	type args struct {
		path string
	}

	usr, _ := user.Current()
	dir := usr.HomeDir

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "A tidle",
			args: args{path: "~"},
			want: dir,
		},
		{
			name: "A tidle path",
			args: args{"~/some"},
			want: dir + "/some",
		},
		{
			name: "An absolute path",
			args: args{"/home/tom/some"},
			want: "/home/tom/some",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandPath(tt.args.path); got != tt.want {
				t.Errorf("ExpandPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
