package main

import (
    "fmt"
    "sort"
)

/*
https://leetcode.cn/problems/4sum/
*/
//n3
func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    for i := 0; i < len(nums)-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            head, tail := j+1, len(nums)-1
            for tail > head {
                if head > j+1 && nums[head] == nums[head-1] {
                    head++
                    continue
                }
                if nums[i]+nums[j]+nums[head]+nums[tail] > target {
                    tail--

                } else if nums[i]+nums[j]+nums[head]+nums[tail] < target {
                    head++

                } else if nums[i]+nums[j]+nums[head]+nums[tail] == target {
                    res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
                    head++
                    tail--

                }

            }
        }
    }
    return res
}
func main() {
    //nums := []int{1, 0, -1, 0, -2, 2}
    nums := []int{2, 2, 2, 2, 2}
    target := 8
    fmt.Println(fourSum(nums, target))
}
