package main

import (
	util "github.com/hktalent/go-utils"
)

func main() {
	util.DoInitAll()
	go util.DoRunning()
	out := util.GetSimpleInput()
	for x := range out {
		if "" != *x {
			func(s1 string) {
				util.DefaultPool.Submit(func() {
					if m1 := util.GetUrlInfo(s1); nil != m1 {
						m1.Set("tools", "httpHeader")
						m1.Set("tags", "httpHeader,http")
						util.PushLog(&m1)
					}
				})
			}(*x)
		} else {
			break
		}
	}
	util.Wg.Wait()
	util.CloseAll()
	util.CloseLogBigDb()
}
