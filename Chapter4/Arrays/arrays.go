package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)


var pc [256]byte


func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("Y"))

	fmt.Println(bitsDiff(&c1,&c2))
	fmt.Println(compare(&c1,&c2))
}

func compare(aptr *[32]byte, bptr *[32]byte) int {
	count := 0

	for i := range aptr {
		a := PopCount(aptr,i)
		b := PopCount(bptr,i)

		diff := math.Abs(float64(a-b))
		count+= int(diff)
	}
	return count
}

func PopCount(ptr *[32]byte,x int) int {
	return int(ptr[byte(x>>(0*8))]+
	ptr[byte(x>>(1*8))]+
	ptr[byte(x>>(2*8))]+
	ptr[byte(x>>(3*8))]+
	ptr[byte(x>>(4*8))]+
	ptr[byte(x>>(5*8))]+
	ptr[byte(x>>(6*8))]+
	ptr[byte(x>>(7*8))])
}

func bitsDiff(b1, b2 *[sha256.Size]byte) int {
    var sum int
    for i := 0; i < sha256.Size; i++ {
        sum += int(pc[b1[i]^b2[i]])
    }
    return sum
}