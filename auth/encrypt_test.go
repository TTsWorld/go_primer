package main

import (
	"fmt"
	"testing"

	"github.com/marspere/goencrypt"
)

func TestA(t *testing.T) {
	Encrypt()

}

func Encrypt() (result string) {
	var err error
	cipher, err := goencrypt.NewAESCipher([]byte("sugosugosugosugo"), []byte("0123456789asdfgh"), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err = cipher.AESEncrypt([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	cipher2, err := goencrypt.NewAESCipher([]byte("sugosugosugosugo"), []byte("0123456789asdfgh"), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		fmt.Println(err)
		return
	}
	plainText, err := cipher2.AESDecrypt(result)
	fmt.Println(plainText)
	return
}
