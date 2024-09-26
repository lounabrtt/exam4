package main

import (
    "fmt"
)

func Ft_missing(nums []int) int {
    n := len(nums)
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    for _, num := range nums {
        actualSum += num
    }
    return expectedSum - actualSum
}

func main() {
    fmt.Println(Ft_missing([]int{3, 1, 2}))                   // resultat : 0
    fmt.Println(Ft_missing([]int{0, 1}))                      // resultat : 2
    fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8
}
