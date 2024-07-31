package main

import (
	"os"
	"piscine"
)

func main() {
	if len(os.Args) > 0 {
		piscine.PrintStr(os.Args[0][2:len(os.Args[0])] + "\n")
	}
}
