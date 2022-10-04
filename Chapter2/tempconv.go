package main

import (
	"fmt"

	"./conv"
)

func main() {
	fmt.Printf("0K = %gºC\n", conv.KtoC(0))
	fmt.Printf("0ºC = %gK\n", conv.CtoK(0))
	fmt.Printf("0ºF = %gK\n", conv.FtoK(0))
}