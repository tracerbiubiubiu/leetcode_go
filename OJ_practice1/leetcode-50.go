package main

import "fmt"

/*删除有序数组中的重复项
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
//没有修改数组
func removeDuplicates1(nums []int) int {
    res := 1
    if len(nums) == 0 {
        return 0
    }
    l := 0
    p := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[l] {
            nums[p] = nums[i]
            l = i
            p++
            res++
        }
    }
    return res
}
func main() {
    nums := []int{1, 1, 2}
    fmt.Println(removeDuplicates1(nums))
    fmt.Println(nums)
}

//原地修改数组
//func removeDuplicates(nums []int) int {
//    n := len(nums)
//    if n == 0 {
//        return 0
//    }
//    slow := 1
//    for fast := 1; fast < n; fast++ {
//        if nums[fast] != nums[fast-1] {
//            nums[slow] = nums[fast]
//            slow++
//        }
//    }
//    return slow
//}
