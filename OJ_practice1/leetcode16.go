package main

import (
    "fmt"
    "math"
    "sort"
)

/*
https://leetcode.cn/problems/3sum-closest/
*/

func doDivide(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
func threeSumClosest(nums []int, target int) int {
    res := math.MaxInt32
    res1 := 0
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {

        head, tail := i+1, len(nums)-1
        for tail > head {
            tmp := doDivide(target, nums[head]+nums[tail]+nums[i])
            if tmp < res {
                res = tmp
                res1 = nums[head] + nums[tail] + nums[i]
            }
            if tail > head+1 && nums[head]+nums[tail]+nums[i]-target > 0 {
                tail--
                continue
            } else if head+1 < tail && nums[head]+nums[tail]+nums[i]-target < 0 {
                head++
                continue
            } else if nums[head]+nums[tail]+nums[i]-target == 0 {
                return target
            }
            break

        }
    }
    return res1
}
func main() {
    nums := []int{-1, 2, 1, -4}
    target := 1
    fmt.Println(threeSumClosest(nums, target))
}
