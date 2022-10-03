package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++{
		fmt.Println(i, (os.Args[i]))
	}
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

