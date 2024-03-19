package main

import (
	"GoSecretsAPI/encrypt"
	"crypto/md5"
)

func main() {
	hasher = md5.New()
	hasher.Write([]byte(key))
	encrypt.Encrypt("dfghjk")
}
