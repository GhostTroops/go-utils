package go_utils

import (
	"bytes"
	"encoding/json"
	"sync"
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

// 避免重复，并设置标记
func CheckNoRepeat4Onece(m *sync.Map, k interface{}) bool {
	if _, ok := m.Load(k); ok {
		return true
	}
	m.Store(k, true)
	return false
}
