package hash

import (
	"crypto/sha256"
	"fmt"
	)

func Sha256Test() {
	s := "세종사이버대 정보보호학과 암호이론 Hash 테스트 예제"

	h := sha256.New()

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}
