package main

import (
	"github.com/GoSecretsAPI/secret"
	"fmt"
)

func main() {
	v := secret.Memory("frhivuwefher")
	err := v.Set("anexamplekey", "a value of strings")
	if err != nil {
		panic(err)
	}
	plaintext, err := v.Get("anexamplekey")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plaintext)
}
