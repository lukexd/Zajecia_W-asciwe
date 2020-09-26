package main

import (
	"fmt"
	"github/Ko4s/goCourse/class2"
)

func main() {
	filePath := "class2/test.txt"
	r := class2.OpenFile(filePath)
	fmt.Println(r)
}
