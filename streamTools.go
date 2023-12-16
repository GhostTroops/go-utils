package go_utils

import (
	"bytes"
	"io"
	"log"
)

// 读取流，逐行 cbk
func ReadStream4Line(r1 io.Reader, cbk func(*string)) {
	var data []byte
	var td = make([]byte, 10240)
	var lSp = []byte("\n")
	for {
		if i, err := r1.Read(td); err == nil {
			if 0 < i {
				if n := bytes.Index(td, lSp); -1 < n {
					data = append(data, td[0:n]...)
					line := string(data)
					data = []byte{}
					cbk(&line)
					if n+1 < i {
						data = td[n+1 : i]
					}
				} else {
					data = append(data, td[0:i]...)
				}
			}
		} else if err == io.EOF {
			break
		} else {
			log.Println(err)
		}
	}
	if nil != data && 0 < len(data) {
		line := string(data)
		cbk(&line)
	}
}
