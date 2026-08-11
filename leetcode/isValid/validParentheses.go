package main 

import (
	"fmt"
)

func isValid(s string) bool {
	stack := make([]byte, 0)
	pairs := map[byte]byte{']': '[', '}': '{', ')': '('}

	for i := 0; i < len(s); i++ {
		sp := len(stack) - 1
		switch c := s[i]; c {
		case '[', '{', '(':
			stack = append(stack, c)
		case ']', '}', ')':
			if len(stack) == 0 || stack[sp] != pairs[c] {
				return false
			}
			stack = stack[:sp]
		}
	}

	return len(stack) == 0
}

func main()  {
	s := "(()))"
	fmt.Println(isValidV2(s))
}

