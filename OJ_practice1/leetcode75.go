package main

import "fmt"

/*
https://leetcode.cn/problems/sort-colors/
*/
//单指针比较2次
func sortColors(nums []int) {
    p0 := 0
    for i := range nums {
        if nums[i] == 0 {
            nums[p0], nums[i] = nums[i], nums[p0]
            p0++
        }
    }

    for i := range nums {
        if nums[i] == 1 {
            nums[i], nums[p0] = nums[p0], nums[i]
            p0++
        }
    }
}
func main() {
    a := []int{2, 0, 2, 1, 1, 0}
    sortColors(a)
    fmt.Println(a)
}
