package main

import (
	go_utils "github.com/hktalent/go-utils"
	"log"
)

func main() {
	n := go_utils.GetValAsInt64("UploadFileMaxSize", 888)
	log.Println(n)

	go_utils.Wg.Wait()
	go_utils.CloseAll()
}
