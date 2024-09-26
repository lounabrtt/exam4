package main

import (
    "fmt"
)

func Ft_min_window(s string, t string) string {
    if len(s) == 0 || len(t) == 0 {
        return ""
    }

    targetCount := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        targetCount[t[i]]++
    }

    left, right, start, minLen, formed := 0, 0, 0, len(s)+1, 0
    windowCount := make(map[byte]int)

    for right < len(s) {
        windowCount[s[right]]++

        if windowCount[s[right]] == targetCount[s[right]] {
            formed++
        }

        for left <= right && formed == len(targetCount) {
            if right-left+1 < minLen {
                start = left
                minLen = right - left + 1
            }

            windowCount[s[left]]--
            if windowCount[s[left]] < targetCount[s[left]] {
                formed--
            }
            left++
        }
        right++
    }

    if minLen == len(s)+1 {
        return ""
    }

    return s[start : start+minLen]
}

func main() {
    fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
    fmt.Println(Ft_min_window("a", "aa"))              // resultat : ""
}
