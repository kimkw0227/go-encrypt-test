package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func RSAAsymTest() {
	privateKey, err := rsa.GenerateKey(rand.Reader,1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	s := `이것은 세종사이버대 정보보호대학원 
		암호이론수업 과제를 위한 대칭/비대칭키 암호화
		성능 테스트 예제임`

	ciphertext, err := rsa.EncryptPKCS1v15(
			rand.Reader,
			publicKey,
			[]byte(s),
		)

	fmt.Printf("%x\n",ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15(
			rand.Reader,
			privateKey,
			ciphertext,
		)

	fmt.Println(string(plaintext))
}
