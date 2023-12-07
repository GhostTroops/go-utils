package go_utils

import (
	"bytes"
	"encoding/json"
)

// 格式化 map 并返回 str
func Map2FormatStr(m *map[string]interface{}) string {
	if data, err := Json.Marshal(m); nil == err {
		var out bytes.Buffer
		if nil == json.Indent(&out, data, "", "\t") {
			return out.String()
		}
	}
	return ""
}
