package main

import "fmt"

/*
https://leetcode.cn/problems/ZVAVXX/
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
    count := 0
    i := 0
    cur := 1

    for j := range nums {
        cur *= nums[j]
        for ; i <= j && cur >= k; i++ {
            cur /= nums[i]
        }
        count += j - i + 1
    }

    return count
}
func main() {
    nums := []int{10, 5, 2, 6}
    k := 100
    fmt.Println(numSubarrayProductLessThanK(nums, k))
}

//func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
//    prod, i := 1, 0
//    for j, num := range nums {
//        prod *= num
//        for ; i <= j && prod >= k; i++ {
//            prod /= nums[i]
//        }
//        ans += j - i + 1
//    }
//    return
//}
