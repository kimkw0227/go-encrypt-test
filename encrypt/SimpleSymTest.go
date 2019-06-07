package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func SimpleSymTest() {
	text := []byte("My Super Secret Code stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm,err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte,gcm.NonceSize())

	if _,err = io.ReadFull(rand.Reader,nonce); err != nil {
		fmt.Println(err)
	}

	fmt.Println(gcm.Seal(nonce,nonce,text,nil))

}