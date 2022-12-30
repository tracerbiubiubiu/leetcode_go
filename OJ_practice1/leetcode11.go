package main

/*
https://leetcode.cn/problems/container-with-most-water/
*/
func maxArea(height []int) int {
    res := 0
    left, right := 0, len(height)-1
    calcArea := func(left, right int) int {
        y := minOf2(height[left], height[right])
        x := right - left
        return x * y
    }
    for left < right {
        res = maxOf2(res, calcArea(left, right))
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return res
}
func minOf2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}
