package go_utils

import "testing"

func TestGetAddr(t *testing.T) {
	type args struct {
		s      string
		szType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test tcp",
			args: args{":0", "tcp"},
			want: "[::]:54878",
		},
		{
			name: "test udp",
			args: args{":0", "udp"},
			want: "[::]:54878",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAddr(tt.args.s, tt.args.szType); got != tt.want {
				t.Errorf("GetAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
