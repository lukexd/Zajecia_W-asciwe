package main

import (
	"fmt"
	"github/Ko4s/goCourse/class2"
)

func main() {

	//class2 examples
	//Czytanie plików
	fileName := "./class2/test.txt"
	class2.ReadFile(fileName)

	//Hashowanie
	word2Hash := "Moje sekretne haslo uzywane na kazdej stronie"
	hash := class2.HashExample(word2Hash)
	fmt.Println("")
	fmt.Printf("Hashed value: %x", hash)

}
