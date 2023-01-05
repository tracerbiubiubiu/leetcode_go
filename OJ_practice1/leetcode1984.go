package main

import (
    "fmt"
    "math"
    "sort"
)

/*
https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
*/
func minimumDifference(nums []int, k int) int {
    if len(nums) == 1 {
        return 0
    }
    res := math.MaxInt32
    sort.Ints(nums)
    for i := 0; i <= len(nums)-k; i++ {
        j := i + k - 1
        res = minOf2(res, nums[j]-nums[i])
    }
    return res
}
func minOf2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
func main() {
    a := []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
    k := 6
    fmt.Println(minimumDifference(a, k))
}
