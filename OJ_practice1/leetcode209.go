package main

import "math"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。


*/
func minSubArrayLen(target int, nums []int) int {
    start, end := 0, 0
    sum := 0
    res := math.MaxInt
    for i, _ := range nums {
        sum += nums[i]
        for sum >= target {
            res = minOf2(res, end-start+1)
            sum -= nums[start]
            start++
        }
        end++
    }
    if res == math.MaxInt {
        return 0
    }
    return res
}
func minOf2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
