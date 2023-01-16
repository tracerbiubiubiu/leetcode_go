package main

import (
    "fmt"
    "strconv"
)

/*
https://leetcode.cn/problems/optimal-division/
*/
func optimalDivision(nums []int) string {
    if len(nums) == 1 {
        return strconv.Itoa(nums[0])
    }
    if len(nums) == 2 {
        return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
    }
    top := nums[0]
    res := strconv.Itoa(top)
    res += "/("
    for i := range nums {
        if i == 0 {
            continue
        }
        if i == len(nums)-1 {
            res += strconv.Itoa(nums[i]) + ")"
            continue
        }
        res += strconv.Itoa(nums[i]) + "/"
    }
    return res

}
func main() {
    nums := []int{1000, 100, 10, 2}
    fmt.Println(optimalDivision(nums))
}

//分母越小，分子越大，
/*


分子反过来，要取最小的为分母，后面不用判断，不加（）直接除==>分子区第一个 其他的都放在分母不用（）
*/
