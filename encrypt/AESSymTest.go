package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte,aes.BlockSize - mod)
		plaintext = append(plaintext,padding...)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b,iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:],plaintext)

	return ciphertext
}

func decrypt(b cipher.Block,ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수여야함.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte,len(ciphertext))
	mode := cipher.NewCBCDecrypter(b,iv)

	mode.CryptBlocks(plaintext,ciphertext)

	return plaintext
}

func AESSymTest() {
	key := "Hello Key 654321"
	s := `이것은 세종사이버대 정보보호대학원 
		암호이론수업 과제를 위한 대칭/비대칭키 암호화
		성능 테스트 예제임`

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block,ciphertext)
	fmt.Println(string(plaintext))

}