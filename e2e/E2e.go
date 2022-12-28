package main

import (
	"github.com/matrix-org/gomatrix"
	"log"
)

/*
https://matrix.org/.well-known/matrix/server
matrix-federation.matrix.org

curl -XPOST -d '{"type":"m.login.password", "user":"example", "password":"wordpass"}' \
 "https://matrix.org/_matrix/client/api/v1/login"

*/
func main() {
	// 创建一个新的Matrix客户端
	client, err := gomatrix.NewClient("https://matrix-client.matrix.org", "", "")
	if err != nil {
		log.Fatal(err)
	}

	// 登录到服务器
	resp, err := client.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     "s1pwn",
		Password: "example_password",
	})
	if err != nil {
		log.Fatal(err)
	}
	/*
		{
		    "access_token": "QGV4YW1wbGU6bG9jYWxob3N0.vRDLTgxefmKWQEtgGd",
		    "home_server": "localhost",
		    "user_id": "@example:localhost"
		}
	*/
	log.Printf("%+v\n", resp)
	// curl -XGET "http://localhost:8008/_matrix/client/r0/sync?access_token=YOUR_ACCESS_TOKEN"
	/*
			{
			    "next_batch": "s72595_4483_1934",
			    "rooms": {
			        "join": {
			            "!726s6s6q:example.com": {
			                "state": {
			                    "events": [
			                        ...
			                    ]
			                },
			                "timeline": {
			                    "events": [
			                        ...
			                    ]
			                }
			            },
			            ...
			        }
			    }
			}

		//  use the "next_batch" token
		curl -XGET "http://localhost:8008/_matrix/client/r0/sync?access_token=YOUR_ACCESS_TOKEN&since=s72595_4483_1934"

	*/
	// 发送消息到给定的房间ID
	_, err = client.SendText("!room_id:matrix.org", "Hello, world!")
	if err != nil {
		log.Fatal(err)
	}
}
