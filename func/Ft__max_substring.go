package main

import (
    "fmt"
)

func Ft_max_substring(s string) int {
    charIndex := make(map[byte]int)
    maxLen, start := 0, 0

    for i := 0; i < len(s); i++ {
        if lastIndex, found := charIndex[s[i]]; found && lastIndex >= start {
            start = lastIndex + 1
        }
        charIndex[s[i]] = i
        maxLen = max(maxLen, i-start+1)
    }

    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
    fmt.Println(Ft_max_substring("bbbbb"))    // resultat : 1
}
