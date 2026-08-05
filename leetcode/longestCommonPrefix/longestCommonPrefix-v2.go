package main

import (
	"fmt"
)

func longestCommonPrefixv2(strs []string) string {
	l := len(strs)
	prefix := make([]byte, 0, 64)
	j := 0

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; ; i++ {
		i2 := i % l
		if len(strs[i2]) == 0 {
			return ""
		}
        if len(prefix) <= j {
            prefix = append(prefix, 0)
        }
		if j == len(strs[i2]) {
			return string(prefix[:j])
		}
		if i2 == 0 {
			prefix[j] = strs[i2][j]
		}
		if strs[i2][j] != prefix[j] {
			return string(prefix[:j])
		}
		if i2 == l - 1 {
			j++
		}
	}
}

func main()  {
	test := []string {"flower","flow","flight"}
	fmt.Println(len(test))
	fmt.Println(longestCommonPrefixv2(test))
}
