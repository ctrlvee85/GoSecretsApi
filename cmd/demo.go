package main

import (
	"GoSecretsAPI/secret"
	"fmt"
)

func main() {
	v := secret.Memory("hhhhh")
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
