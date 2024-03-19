package encrypt

import (
	"encoding/hex"
	"crypto/cipher"
	"errors"
	"crypto/md5"
	"crypto/aes"
	"fmt"
	"io"
	"crypto/rand"
)

func Encrypt(key, plaintext string) (string, error) {

	// key to encode secret string
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}
	//read key
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader,iv); err != nil {
		return "", err
	}

	// key destination
	stream := cipher.NewCFBEncrypter(block,iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", cipherText), nil

}

func Decrypt(key, cipherHex string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	cipherText,err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText,cipherText)


	return string(cipherText), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}