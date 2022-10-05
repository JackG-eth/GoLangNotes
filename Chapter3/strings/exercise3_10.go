package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(commaIterative("Hello!!World!!"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return(comma(s[:n-3])+ "," + s[n-3:])
}

func commaIterative(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
    d, t := len(s)%3, len(s)/3

    if d != 0 {
        buf.WriteString(s[:d])
        buf.WriteByte(',')
    }

    for i := 0; i < t; i++ {
        buf.WriteString(s[i*3+d:d+i*3+3])
        if i != t-1 {
            buf.WriteByte(',')
        }
    }
    return buf.String()
}

// s[:n-3]) n == 10 so this would print all the charcters up to 7
//  s[n-3:]n == 10 so this would print all the charcters past 3 etc