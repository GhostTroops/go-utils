package main

import (
	util "github.com/hktalent/go-utils"
	"log"
)

func main() {
	n := util.GetValAsInt64("UploadFileMaxSize", 888)
	log.Println(n)

	util.Wg.Wait()
	util.CloseAll()
}
