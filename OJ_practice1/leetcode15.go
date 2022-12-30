package main

import (
    "fmt"
    "sort"
)

/*

 */
//去重用map
//头尾指针

//func threeSum(nums []int) (res [][]int) {
//
//    sort.Ints(nums)
//    for i := 0; i < len(nums)-2; i++ {
//        head, tail := 1, len(nums)-1
//        n1, n2, n3 := nums[i], nums[head], nums[tail]
//        if nums[i] > 0 {
//            break
//        }
//        if i >= 1 && nums[i-1] == nums[i] {
//            continue
//        }
//        for tail > head {
//            if n2+n3-n1 > 0 {
//                tail--
//            } else if n2+n3-n1 < 0 {
//                head++
//            } else {
//                res = append(res, []int{n1, n2, n3})
//                for tail > head && n2 == nums[head] {
//                    head++
//                }
//                for tail > head && nums[tail] == nums[tail] {
//                    tail--
//                }
//            }
//        }
//    }
//    return res
//}
func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    if nums[0] > 0 || nums[len(nums)-1] < 0 {
        return res
    }
    for i := 0; i < len(nums)-2; i++ {
        head, tail := i+1, len(nums)-1
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for head < tail {
            //if head > i+1 && nums[head] == nums[head-1] {
            //    head++
            //    continue
            //}
            if tail+1 <= len(nums)-1 && nums[tail] == nums[tail+1] {
                tail--
                continue
            }
            if nums[i]+nums[head]+nums[tail] > 0 {
                tail--

            } else if nums[i]+nums[head]+nums[tail] < 0 {
                head++

            } else {
                res = append(res, []int{nums[i], nums[head], nums[tail]})
                tail--
                head++
            }

        }
    }
    return res
}
func main() {
    nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
    //nums := []int{-1, 0, 1, 2, -1, -4}
    fmt.Println(threeSum(nums))
}
