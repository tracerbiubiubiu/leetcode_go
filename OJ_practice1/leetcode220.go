package main

import "fmt"

func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

//差几个用例未通过的笨方法
//func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
//    for i := 0; i <= len(nums)-indexDiff; i++ {
//        j := i + indexDiff
//        if j >= len(nums) {
//            j = len(nums) - 1
//        }
//        for j > i {
//            if abs(nums[i], nums[j]) <= valueDiff {
//                return true
//            }
//            j--
//        }
//
//    }
//    return false
//}

//func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
//
//}
func main() {
    nums := []int{2, 2}
    indexDiff := 2
    valueDiff := 0
    fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
}
