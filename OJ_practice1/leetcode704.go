package main

/*
https://leetcode.cn/problems/binary-search/description/
*/
func search(nums []int, target int) int {
    length := len(nums)
    mid := (length - 1) / 2
    left, right := 0, length-1
    for left <= right {
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            return mid
        }
        mid = left + (right-left)/2
    }
    return -1
}
