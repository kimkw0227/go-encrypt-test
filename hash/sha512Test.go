package hash

import (
	"crypto/sha512"
	"fmt"
	)

func Sha512Test() {
	s := "세종사이버대 정보보호학과 암호이론 Hash 테스트 예제"

	h := sha512.New()

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}
