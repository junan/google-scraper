package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt() {
	password := "asd123asd"
	bcrypt.GenerateFromPassword([]byte(password), 14)
}

func main() {
	fmt.Println("Hello World")
}
