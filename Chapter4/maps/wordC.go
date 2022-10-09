package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	
    counts := make(map[string]int) // counts of Unicode characters
    input := bufio.NewScanner(os.Stdin)

    // set SplitFunc with ScanWords instead of default ScanLines
    // ScanWords just return space-separated words
    input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Println("--------Output-------")
	for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }

}