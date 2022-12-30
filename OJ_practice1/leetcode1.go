package main

/*

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//暴力 双指针
func twoSum1(nums []int, target int) []int {
    if len(nums) == 0 {
        return nil
    }
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            if nums[j] == target-nums[i] {
                return []int{i, j}
            }
        }
    }
    return nil
}

//哈希表
//func twoSum2(nums []int, target int) []int {
//    hashTable := map[int]int{}
//    for i, x := range nums {
//        if p, ok := hashTable[target-x]; ok {
//            return []int{p, i}
//        }
//        hashTable[x] = i
//    }
//    return nil
//}
//
func twoSum2(nums []int, target int) []int {
    tmpMat := make(map[int]int, 0)

    for i, v := range nums {
        if _, ok := tmpMat[target-v]; ok {
            return []int{i, tmpMat[target-v]}
        } else {
            tmpMat[v] = i
        }
    }
    return []int{}
}
