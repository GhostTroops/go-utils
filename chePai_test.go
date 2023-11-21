package go_utils

import (
	"fmt"
	"log"
	"testing"
)

func TestWhereCar(t *testing.T) {
	log.Println(WhereCar("WJ210034"))
	log.Println(WhereCar("VV10034"))
	log.Println(WhereCar("川G0034"))
	log.Println(WhereCar("川A0034"))
	log.Println(WhereCar("使2170034"))
	log.Println(WhereCar("台C70034"))
	log.Println(WhereCar("民航034"))
	fmt.Println(Fanyi4YoudaoPars(`こんにちは`, "ja"))
	fmt.Println(Fanyi4YoudaoPars(`안녕하세요`, "ko"))
}
