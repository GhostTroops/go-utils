package go_utils

import (
	"testing"
)

func TestTransInt64To62(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试62进制转换 1", args{889533}, "cMvq"},
		{"测试62进制转换 2", args{11}, "j"},
		{"测试62进制转换 3", args{0}, "a"},
		{"测试62进制转换 4", args{10}, "i"},
		{"测试62进制转换 5", args{62}, "ba"},
		{"测试62进制转换 6", args{61}, "Z"},
		{"测试62进制转换 7", args{0xff}, "d8"},
	}
	var aR []string
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransInt64To62(tt.args.id); got != tt.want {
				t.Errorf("TransInt64To62() = %v, want %v", got, tt.want)
			} else {
				aR = append(aR, got)
			}
		})
	}

	type args2 struct {
		id string
	}
	tests1 := []struct {
		name string
		args args2
		want int64
	}{
		{"fz 1", args2{aR[0]}, 889533},
		{"fz 2", args2{aR[1]}, 11},
		{"fz 3", args2{aR[2]}, 0},
		{"fz 4", args2{aR[3]}, 10},
		{"fz 5", args2{aR[4]}, 62},
		{"fz 6", args2{aR[5]}, 61},
		{"fz 7", args2{aR[6]}, 255},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trans62ToInt64(tt.args.id); got != tt.want {
				t.Errorf("TransInt64To62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRadomTemplate(t *testing.T) {
	type args struct {
		id int64
	}
	szTp := GetRadomTemplate()
	t.Log("template: ", szTp)
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试62进制转换 1", args{889533}, "cMvq"},
		{"测试62进制转换 2", args{11}, "j"},
		{"测试62进制转换 3", args{0}, "a"},
		{"测试62进制转换 4", args{10}, "i"},
		{"测试62进制转换 5", args{62}, "ba"},
		{"测试62进制转换 6", args{61}, "Z"},
		{"测试62进制转换 7", args{0xff}, "d8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransInt64ToN(tt.args.id, szTp); tt.args.id != TransN2Int64(got, szTp) {
				t.Errorf("TransInt64To62() = %v, want %v", got, tt.args.id)
			} else {
				t.Logf("%d <= to => %s is ok", tt.args.id, got)
			}
		})
	}
}
