package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
	} else {
		fileYmkn, err := os.ReadFile(os.Args[1])
		if err != nil {
			return
		}
		fmt.Print(string(fileYmkn))
	}
}
