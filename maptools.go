package go_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

func MergeObjs(i ...interface{}) *map[string]interface{} {
	var m1 = map[string]interface{}{}
	for _, x := range i {
		if data, err := Json.Marshal(x); nil == err {
			Json.Unmarshal(data, &m1)
		}
	}
	return &m1
}

// 格式化 map 并返回 str
func Map2FormatStr(m *map[string]interface{}) string {
	var lk = GetLock("Map2FormatStr").Lock()
	defer lk.Unlock()
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

/*
移除空的、无效的值
*/
func RmNullMap(m *map[string]interface{}) *map[string]interface{} {
	if nil == m {
		return nil
	}
	for k, v := range *m {
		if m1, ok := v.(map[string]interface{}); ok {
			(*m)[k] = RmNullMap(&m1)
			continue
		}
		s1 := fmt.Sprintf("%v", v)
		if nil == v || s1 == "null" || "nil" == s1 || "" == s1 {
			delete(*m, k)
		}
	}
	return m
}
