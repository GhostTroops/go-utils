package go_utils

import (
	"fmt"
	"log"
	"testing"
)

func TestDetectProxy(t *testing.T) {
	fmt.Println(Fanyi4YoudaoPars("love", "en"))
	//type args struct {
	//	szIp   string
	//	szPort string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := DetectProxy(tt.args.szIp, tt.args.szPort); got != tt.want {
	//			t.Errorf("DetectProxy() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func TestCheckHttsProxy(t *testing.T) {
	type args struct {
		szIp   string
		szPort string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "https proxy",
			args: args{"1.0.205.87", "8080"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DetectProxy(tt.args.szIp, tt.args.szPort) {
				log.Println("ok")
			}
		})
	}
}
