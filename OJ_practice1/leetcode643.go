package main

import "fmt"

/*
https://leetcode.cn/problems/maximum-average-subarray-i/
*/
//超时
func findMaxAverage(nums []int, k int) float64 {
    var sum float64
    for _, v := range nums[:k] {
        sum += float64(v)
    }
    maxSum := sum
    for i := 1; i <= len(nums)-k; i++ {
        j := i + k - 1
        sum = sum + float64(nums[j]) - float64(nums[i-1])
        maxSum = maxOf22(maxSum, sum)
    }
    return maxSum / float64(k)
}

func maxOf22(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
func main() {
    // nums := []int{0, 1, 1, 3, 3}
    nums := []int{0, 4, 0, 3, 2}
    k := 1
    fmt.Println(findMaxAverage(nums, k))
}
