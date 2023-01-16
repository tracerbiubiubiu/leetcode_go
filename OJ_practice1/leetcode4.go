package main

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。


*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length := len(nums1) + len(nums2)
    if length%2 == 0 { //length&1==0
        k1 := length / 2
        k2 := length/2 + 1
        return float64(handleArray(nums1, nums2, k1)+handleArray(nums1, nums2, k2)) / 2.0
    } else {
        k := (length + 1) / 2
        return float64(handleArray(nums1, nums2, k))
    }

}

//func getMidNum(nums1, nums2 []int, k int) int {
//    idx1, idx2 := 0, 0
//
//    for {
//        if idx1 == len(nums1) {
//            return nums2[idx2+k-1]
//        }
//        if idx2 == len(nums2) {
//            return nums1[idx1+k-1]
//        }
//        if k == 1 {
//            return min(nums1[idx1], nums2[idx2])
//        }
//
//        half := k / 2
//        newIdx1 := min(len(nums1), idx1+half) - 1
//        newIdx2 := min(len(nums2), idx2+half) - 1
//        if nums1[newIdx1] <= nums2[newIdx2] {
//            k -= newIdx1 + 1 - idx1
//            idx1 = newIdx1 + 1
//        } else {
//            k -= newIdx2 + 1 - idx2
//            idx2 = newIdx2 + 1
//        }
//    }
//}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

//我想的是用递归写
//也是每次裁剪数组，更新k，也是找第k小的数
//比官方解答好像好一些

func handleArray(nums1, nums2 []int, k int) int {
    //idx1,idx2 := 0,0
    for {
        if len(nums1) == 0 {
            return nums2[k-1]
        }
        if len(nums2) == 0 {
            return nums1[k-1]
        }
        if k == 1 {
            return min(nums1[0], nums2[0])
        }
        half := k / 2
        //这里是坑 别忘了
        compareLen1 := min(len(nums1), half)
        compareLen2 := min(len(nums2), half)
        if nums1[compareLen1-1] <= nums2[compareLen2-1] {
            nums1 = nums1[compareLen1:]
            k -= compareLen1
            return handleArray(nums1, nums2, k)
        } else {
            nums2 = nums2[compareLen2:]
            k -= compareLen2
            return handleArray(nums1, nums2, k)
        }
    }
}
