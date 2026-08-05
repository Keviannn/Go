package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	prefix := make([]byte, 0, 64)
	j := 0

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; ; i++ {
		if i >= l {
			i = 0
		}
		if len(strs[i]) == 0 {
			return ""
		}
        if len(prefix) <= j {
            prefix = append(prefix, 0)
        }
		if j == len(strs[i]) {
			return string(prefix[:j])
		}
		if i == 0 {
			prefix[j] = strs[i][j]
		}
		if strs[i][j] != prefix[j] {
			return string(prefix[:j])
		}
		if i == l - 1 {
			j++
		}
	}
}

func main()  {
	test := []string {"flower","flow","flight"}
	fmt.Println(len(test))
	fmt.Println(longestCommonPrefix(test))
}
