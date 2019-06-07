package main

import (
	"fmt"
	"go-encrypt-test/encrypt"
	"go-encrypt-test/hash"
	"go-encrypt-test/sign"
	"os"
	"time"
)
func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("**********************************************")
		fmt.Println("\tHashTest: go run hash")
		fmt.Println("\tEncryptTest: go run encrypt")
		fmt.Println("\tSignatureTest: go run sign")
		fmt.Println("**********************************************")
	} else {
		switch os.Args[1] {
		case "hash":
			hashTest()
			break
		case "encrypt":
			encryptTest()
			break
		case "sign":
			signTest()
			break
		}
	}
}

func hashTest() {
	startTime := time.Now()
	fmt.Println(startTime)

	for i := 0;i < 10000; i++ {
		hash.Md5Test()
	}
	elapsedTime := time.Since(startTime)

	startTime2 := time.Now()
	for i := 0; i < 10000; i++ {
		hash.Sha1Test()
	}
	elapsedTime2 := time.Since(startTime2)

	startTime3 := time.Now()
	for i := 0; i < 10000; i++ {
		hash.Sha256Test()
	}
	elapsedTime3 := time.Since(startTime3)

	startTime4 := time.Now()
	for i := 0; i < 10000; i++ {
		hash.Sha512Test()
	}
	elapsedTime4 := time.Since(startTime4)

	fmt.Printf("Executino Time(md5): %s\n", elapsedTime)
	fmt.Printf("Executino Time(sha1): %s\n", elapsedTime2)
	fmt.Printf("Executino Time(sha256): %s\n", elapsedTime3)
	fmt.Printf("Executino Time(sha512): %s\n", elapsedTime4)
}

func signTest() {
	startTime := time.Now()

	for i := 0;i < 1000; i++ {
		sign.RSASignatureTest()
	}
	elapsedTime := time.Since(startTime)

	startTime2 := time.Now()

	for i := 0;i < 1000; i++ {
		sign.ECDSASignatureTest()
	}
	elapsedTime2 := time.Since(startTime2)

	fmt.Printf("Execution Time(RSA): %s\n", elapsedTime)
	fmt.Printf("Execution Time(ECDSA): %s\n", elapsedTime2)
}

func encryptTest() {
	startTime := time.Now()

	for i := 0;i < 1000; i++ {
		encrypt.AESSymTest()
	}
	elapsedTime := time.Since(startTime)

	startTime2 := time.Now()

	for i := 0;i < 1000; i++ {
		encrypt.RSAAsymTest()
	}
	elapsedTime2 := time.Since(startTime2)

	fmt.Printf("Execution Time(AES): %s\n", elapsedTime)
	fmt.Printf("Execution Time(RSA): %s\n", elapsedTime2)
}