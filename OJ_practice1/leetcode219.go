package main

import "fmt"

/*
https://leetcode.cn/problems/contains-duplicate-ii/
*/
//func containsNearbyDuplicate(nums []int, k int) bool {
//    mp := make(map[int]int)
//    for i := range nums {
//        if _, ok := mp[nums[i]]; ok {
//            if abs(i, mp[nums[i]]) <= k {
//                return true
//            } else {
//                mp[nums[i]] = i
//            }
//        } else {
//            mp[nums[i]] = i
//        }
//    }
//    return false
//}
func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
func main() {
    //nums := []int{1, 2, 3, 1, 2, 3}
    nums := []int{1, 0, 1, 1}
    k := 1
    fmt.Println(containsNearbyDuplicate(nums, k))
}

//试试滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]bool)
    for i := range nums {
        if i > k {
            mp[nums[i-k-1]] = false
        }
        if mp[nums[i]] {
            return true
        }
        mp[nums[i]] = true
    }
    return false
}
