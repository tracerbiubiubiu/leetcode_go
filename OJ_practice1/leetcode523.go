package main

/*
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。

如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。
https://leetcode.cn/problems/continuous-subarray-sum/description/?languageTags=golang
*/
/*
余数一样，两个数相减结果就是k的整数倍
*/
func checkSubarraySum(nums []int, k int) bool {
    m := make(map[int]int, 0)
    m[0] = -1
    sum := 0
    for i := range nums {
        sum += nums[i]
        if k != 0 {
            sum = sum % k
        }
        if v, ok := m[sum]; ok {
            if i-v > 1 { //i-v+1> 2
                return true
            }
        } else {
            m[sum] = i
        }
    }
    return false
}
