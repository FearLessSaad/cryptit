package main

import (
	"fmt"

	"github.com/FearLessSaad/cryptit/decrypt"
	"github.com/FearLessSaad/cryptit/encrypt"
)

func main() {
	planeText := "HashX Private Limited"
	encrypted := encrypt.Nimbus(planeText)
	decrypted := decrypt.Nimbus(encrypted)

	fmt.Println("PlaneText -", planeText)
	fmt.Println("Encrypted -", encrypted)
	fmt.Println("Decrypted -", decrypted)
}
