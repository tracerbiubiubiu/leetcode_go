package main

import "fmt"

/*
https://leetcode.cn/problems/subarrays-with-k-different-integers/
*/
func subarraysWithKDistinct(nums []int, k int) int {
    tmpMap := make(map[int]int)
    res := 0
    start := 0
    for end := range nums {

        tmpMap[nums[end]]++
        if len(tmpMap) < k {
            continue
        }
        for len(tmpMap) > k {
            tmpMap[nums[start]]--

            if tmpMap[nums[start]] == 0 {
                delete(tmpMap, nums[start])
            }
            start++
        }
        res += end - start - k + 2

    }
    return res
}
func main() {
    a := []int{2, 2, 1, 2, 2, 2, 1, 1}
    //a := []int{2, 1, 1, 1, 2}
    k := 2
    //a := []int{1, 2, 1, 2, 4}
    fmt.Println(subarraysWithKDistinct(a, k))
}
