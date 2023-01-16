package main

import (
    "fmt"
    "math"
)

/*
https://leetcode.cn/problems/increasing-triplet-subsequence/description/
*/
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    first, second := 0, math.MaxInt32
    for i := 1; i < len(nums); i++ {
        if second < len(nums)-1 && nums[i] > nums[second] {
            return true
        }
        if nums[i] > nums[first] {
            second = i
        } else {
            first = i
            //此时存在两个待定条件
        }

    }
    return false
}
func main() {
    //nums := []int{2, 1, 5, 0, 4, 6}
    nums := []int{20, 100, 10, 12, 5, 13}
    fmt.Println(increasingTriplet(nums))
}

//另外一种想法 就是分割数组 左边的最小值小于nums[i] 右边的最大值大于i即可
