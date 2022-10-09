package main

import "fmt"

func main() {
	a := []int{0,1,2,3,4,5}
	a = rotateLeft(a,4)
	fmt.Println(a)

	b := []string {"bob","bob","bob","mike","john","john"}
	b = filteDup(b)
	fmt.Println(b)
}

func reverse(s []int)  {
	for i, j:= 0, len(s) -1; i < j; i, j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	for i, j:= 0, len(s) -1; i < j; i, j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeft(s []int, n int) []int {
    n = n % len(s)
    if n > 0 && n < len(s) {
        temp := make([]int, n)
        copy(temp, s[:n])

        copy(s, s[n:])
        copy(s[len(s)-n:], temp)
    }
    return s
}

func filteDup(s []string) []string {
	if len(s) == 0 {
		return s
	}
	i := 0
	for j := 1; j < len(s); j++ {
		if s[i] != s[j]{
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}