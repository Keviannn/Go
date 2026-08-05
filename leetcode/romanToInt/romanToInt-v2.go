
package main

import (
	"fmt"
)

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

	for i, val := range v {
		if p := priority[val]; i + 1 < vlen && p < priority[v[i + 1]] {
			sum -= p 
		} else {
			sum += p 
		}
	}

	return sum
}

func main () {
	fmt.Println(romanToInt("MCMXCIV"))
}
