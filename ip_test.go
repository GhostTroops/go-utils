package go_utils

import (
	"math/big"
	"reflect"
	"testing"
)

func TestIp2Int(t *testing.T) {
	type args struct {
		ip string
	}
	type Test struct {
		name string
		args args
		want *big.Int
	}

	var tests = []Test{
		Test{name: "Ipv4 to Int",
			args: args{ip: "20.36.77.12"},
			want: big.NewInt(281471019666700),
		},
		Test{name: "Ip v6 to Int",
			args: args{ip: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			want: Str2BigInt("42540766452641154071740215577757643572", 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrIp2Int(tt.args.ip)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ip2Int() = %v, want %v", got, tt.want)
			} else {
				s1 := IntToIpv6Str(got)
				if !reflect.DeepEqual(s1, tt.args.ip) {
					t.Errorf("IntToIpv6Str() = %v, want %v", s1, tt.args.ip)
				}
			}
		})
	}
}

func TestAny2Hex(t *testing.T) {
	type args struct {
		n interface{}
	}
	tests1 := []struct {
		name string
		args args
		want string
	}{
		struct {
			name string
			args args
			want string
		}{name: "int to hex", args: args{
			n: Str2BigInt("42540766452641154071740215577757643572", 0),
		}, want: "3432353430373636343532363431313534303731373430323135353737373537363433353732"}}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any2Hex(tt.args.n); got != tt.want {
				t.Errorf("Int2Hex() = %v, want %v", got, tt.want)
			} else {
				x1 := BigInt2Hex(tt.args.n.(*big.Int), 16)
				x2 := "20010db885a3000000008a2e03707334"
				if x1 != x2 {
					t.Errorf("BigInt2Hex() = %v, want %v", x1, x2)
				}
			}
		})
	}
}
