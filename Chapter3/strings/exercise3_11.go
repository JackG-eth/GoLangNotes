package main

import "fmt"

type ByRune []rune

func main()  {
	fmt.Println(anagram("mona lias", "amon lisa"))
}

func anagram(s1 string, s2 string ) bool {
    m := make(map[rune]int)
	// add each character to the map (increment the counter)
    for _, r := range s1 {
		fmt.Println(r)
        m[r]++
    }
    for _, r := range s2 {
		// loop through all the characters in second string
		// if map contains character remove (if it has multiple decrement)
        if i, ok := m[r]; ok {
            if i > 1 {
                m[r]--
            } else {
                delete(m, r)
            }
        } else {
            return false
        }
    }
    return len(m) == 0
}

