package go_utils

import (
	"fmt"
	"github.com/itchyny/gojq"
	"github.com/simonnilsson/ask"
	"reflect"
)

func isPointer(x interface{}) bool {
	return reflect.TypeOf(x).Kind() == reflect.Ptr
}

// for github.com/itchyny/gojq
func GetJQ(source interface{}, path string) interface{} {
	if isPointer(source) {
		source = reflect.ValueOf(source).Elem()
	}
	if ps, err := gojq.Parse(path); nil == err {
		if data := ps.Run(source); nil != data {
			if o, ok := data.Next(); ok {
				return o
			}
		}
	}
	return nil
}

// itchyny/gojq
func GetJQ2Str(source interface{}, path string) string {
	if ps := GetJQ(source, path); nil != ps {
		return fmt.Sprintf("%v", ps)
	}
	return ""
}

func GetJson4Query(source interface{}, path string) interface{} {
	res := ask.For(source, path)
	if nil != res {
		return res.Value()
	}
	return nil
}
