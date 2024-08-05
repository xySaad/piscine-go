package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("File name missing")
	} else if len(os.Args) > 2 {
		println("Too many arguments")
	} else {
		fileYmkn, err := os.ReadFile(os.Args[1])
		if err != nil {
			return
		}
		fmt.Println(string(fileYmkn))
	}
}
