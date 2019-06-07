package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"
)

func ECDSASignatureTest() {
	publicKeyCurve := elliptic.P256()

	privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(publicKeyCurve,rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	message := "세종사이버대 정보보호학과 암호이론 수업 서명 검증 예제"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err = ecdsa.Sign(rand.Reader, privateKey, digest)
	if err != nil {
		fmt.Println(err)
		return
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	verifystatus := ecdsa.Verify(publicKey,digest,r,s)

	if verifystatus != true {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}

}
