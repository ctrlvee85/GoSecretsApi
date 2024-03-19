package main

import (
	"crypto/aes"
	"crypto/md5"
	"fmt"
)

func main() {
	hasher := md5.New()
	key := []byte("a bunch of words and 5484854 and stuff")
	hasher.Write([]byte(key))
	cipherKey := hasher.Sum(nil)
	fmt.Println(len(cipherKey))
	_, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
}
