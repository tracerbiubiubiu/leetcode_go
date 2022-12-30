package main

import "fmt"

/*

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

 

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]
 

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//切片扩容
//回溯法



func backTrace(nums []int) [][]int {
    n := len(nums)
    res := make([][]int,sliceLength(n))

    if n == 0 {
        return nil
    }

    for i := 0;i < n; i++{
        cur := nums[i]

        nums = append(nums[:i],nums[i+1:]...)
         ans := make([]int,0)
        ans = append(ans,cur)
        backTrace(nums)

        ans = []int{}


    }

}

//func permute(nums []int) [][]int{
//    //n := len(nums)
//    res := make([][]int,sliceLength(len(nums)))
//
//    return res
//}
//func backTrace(nums []int,index int,res [][]int) {
//    //结束条件
//    if index == len(nums) - 1 {
//
//    }
//    for i := index; i < len(nums); i++{
//
//        backTrace(nums,index + 1,res)
//    }
//}
//计算生成切片的长度
func sliceLength(n int) int {
    if n == 1 {
        return 1
    }
    return n * sliceLength(n-1)
}
//生成长度另一种方法
func permuteCnt(n int) int {
    ans := 1
    for i:=0; i<n; i++ {
        ans = ans * (i+1)
    }
    return ans
}

//func permute(nums []int) [][]int {
//    res = [][]int{}
//    backTrack(nums,len(nums),[]int{})
//    return res
//}
//
//func backTrack(nums []int,numsLen int,path []int) {
//    if numsLen == 0 {
//        res = [][]int{}
//    }
//    for i := 0;i < numsLen; i++ {
//        cur := nums[i]
//        //已经使用过的
//        path = append(path,cur)
//        //剩下的
//        nums = append(nums[:i],nums[i+1:]...)
//        backTrack(nums,len(nums),path)
//        //回溯的时候切片也要复原，元素位置不能变
//        nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
//        path = path[:len(path)-1]
//    }
//}

func main() {
    //nums := []int{1,2,3}
    //fmt.Println(permute(nums))
    fmt.Println(sliceLength(4))
}

/*

func permute(nums []int) [][]int {
    n := len(nums)
    ans := make([][]int, 0, permuteCnt(n))
    var backtrace func(nums []int, index int)

    backtrace = func(nums []int, index int) {
        if index == n-1 {
            copyNums := make([]int, n)
            copy(copyNums, nums)
            ans = append(ans, copyNums)
        }

        for i:=index; i<n; i++ {
            nums[i], nums[index] = nums[index], nums[i]
            backtrace(nums, index+1)
            nums[i], nums[index] = nums[index], nums[i]
        }
    }

    backtrace(nums, 0)
    return ans
}

func permuteCnt(n int) int {
    ans := 1
    for i:=0; i<n; i++ {
        ans = ans * (i+1)
    }
    return ans
}

*/