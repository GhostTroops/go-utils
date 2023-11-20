package go_utils

import (
	_ "embed"
	"strings"
)

//go:embed chepai.txt
var chepai string
var ChePm = map[string]string{}

func init() {
	a := strings.Split(chepai, "\n")
	for _, x := range a {
		j := strings.Split(x, "\t")
		if 2 <= len(j) {
			s11 := strings.Join(j[1:], "")
			j[0] = strings.ToLower(j[0])
			ChePm[j[0]] = s11
			ChePm[s11] = j[0]
		}
	}
}

/*
车牌归宿查询：
WhereCar("川G")
WhereCar("黄石市")
*/
func WhereCar(s string) string {
	s = strings.ToLower(s)
	return ChePm[s]
}
