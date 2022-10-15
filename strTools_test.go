package go_utils

import "testing"

func TestTransInt64To62(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试62进制转换 1", args{889533}, "3Jpj"},
		{"测试62进制转换 2", args{11}, "b"},
		{"测试62进制转换 3", args{0}, "0"},
		{"测试62进制转换 4", args{10}, "a"},
		{"测试62进制转换 5", args{62}, "10"},
		{"测试62进制转换 6", args{61}, "Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransInt64To62(tt.args.id); got != tt.want {
				t.Errorf("TransInt64To62() = %v, want %v", got, tt.want)
			}
		})
	}
}
