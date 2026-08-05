package main

import (
	"fmt"
)
/*
var priority = map[byte]int {
	'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
	v := []byte(s)
	vlen := len(v)
	var sum int

	for i := 0; i < vlen; i++ {
		val := v[i]
		if i + 1 < vlen {
			if p1, p2 := priority[val], priority[v[i + 1]]; p1 < p2 {
				sum += p2 - p1
				i++
			} else {
				sum += p1
			}
		} else {
			sum += priority[v[vlen - 1]]
		}
	}

	return sum
}
*/

func main () {
	fmt.Println(romanToInt("MCMXCIV"))
}
