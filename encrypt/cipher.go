package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypt(key []byte) {

	// key to encrypt
	hasher := md5.New()
	hasher.Write([]byte(key))
	key, _ := hex.DecodeString()

}
