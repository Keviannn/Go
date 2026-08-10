package main

import (
	"fmt"
)

func longestCommonPrefixv3(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    first := strs[0]
    for i := 0; i < len(first); i++ {
        c := first[i]
        for _, s := range strs[1:] {
            if i >= len(s) || s[i] != c {
                return first[:i]
            }
        }
    }
    return first
}

func main()  {
	test := []string {"flower","flow","flight"}
	fmt.Println(len(test))
	fmt.Println(longestCommonPrefixv3(test))
}

