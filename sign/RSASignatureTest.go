package sign

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func RSASignatureTest() {
	privateKey, err := rsa.GenerateKey(rand.Reader,1024)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	message := "세종사이버대 정보보호학과 암호이론 수업 서명 검증 예제"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15(
			rand.Reader,
			privateKey,
			h1,
			digest,
		)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(
			publicKey,
			h2,
			digest,
			signature,
		)

	if err != nil {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}

}
