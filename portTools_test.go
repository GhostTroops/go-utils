package go_utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetAddr(t *testing.T) {
	var SkipDnsIp = regexp.MustCompile(`\d+-\d+-\d+-\d+`)
	fmt.Print(SkipDnsIp.MatchString("https://ec2-13-114-147-108.ap-northeast-1.compute.amazonaws.com"))
	//type args struct {
	//	s      string
	//	szType string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	{
	//		name: "test tcp",
	//		args: args{":0", "tcp"},
	//		want: "[::]:54878",
	//	},
	//	{
	//		name: "test udp",
	//		args: args{":0", "udp"},
	//		want: "[::]:54878",
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := GetAddr(tt.args.s, tt.args.szType); got != tt.want {
	//			t.Errorf("GetAddr() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
