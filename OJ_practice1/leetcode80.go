package main

import "fmt"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。
*/
/*
通用解法
为了让解法更具有一般性，我们将原问题的「保留 2 位」修改为「保留 k 位」。

对于此类问题，我们应该进行如下考虑：

由于是保留 k 个相同数字，对于前 k 个数字，我们可以直接保留
对于后面的任意数字，能够保留的前提是：与当前写入的位置前面的第 k 个元素进行比较，不相同则保留

作者：宫水三叶
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/702970/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
//func removeDuplicates(nums []int) int {
//    process := func(k int) int {
//        l := 0
//        p := 1
//        res := 1
//        for i := 1; i < len(nums); i++ {
//            if nums[i] != nums[l] {
//                nums[p] = nums[i]
//                l = i
//                p++
//                res++
//                continue
//            }
//            if i-l < k {
//                nums[p] = nums[i]
//                res++
//                p++
//                continue
//            }
//        }
//        return res
//    }
//    return process(2)
//}

func main() {
    //nums := []int{1, 1, 1, 2, 2, 3}
    //nums := []int{1, 2, 2, 2}
    nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
    fmt.Println(removeDuplicates(nums))
}

/*
func removeDuplicates(nums []int) int {
	var process func(k int) int
	process = func(k int) int {
		u := 0
		for _, v := range nums {
			if u < k || nums[u-k] != v {
				nums[u] = v
				u++
			}
		}
		return u
	}
	return process(2)
}

作者：宫水三叶
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/solutions/702970/gong-shui-san-xie-guan-yu-shan-chu-you-x-glnq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func removeDuplicates(nums []int) int {
    process := func(k int) int {
        res := 0
        w := 0   //写指针
        cur := 0 //当前元素
        //k值 v下标 次数 map不够用 得用map[int]struct{}
        //tmp := make(map[int]int)
        for i := 0; i < len(nums); i++ { //i为读指针
            if nums[i] == nums[cur] {
                if i-cur+1 > k {
                    continue
                }
            } else {
                cur = i
            }
            nums[w] = nums[i]
            res++
            w++

        }
        //fmt.Println(nums)
        return res
    }
    return process(2)
}
