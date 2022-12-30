package main

import "fmt"

/*
那我们要完成一个 a ? a ? a = 0 的运算，是不是其实就是让其二进制的每一位数都相加，最后再对 3 进行一个取模的过程呢？（一样，如果要定义一个 a ? a ? a ? a = 0 的运算，那就最后对 4 进行取模就可以了）

作者：小浩算法
链接：https://leetcode.cn/problems/single-number-ii/solutions/337396/duo-jie-fa-wei-yun-suan-shu-xue-fang-shi-hashmap-b/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

//go只能过一部分是为啥
func singleNumber3(nums []int) int {
    res := 0
    for i := 0; i < 64; i++ {
        //初始化每一位1的个数为0
        number := 0
        for _, k := range nums {
            //通过右移i位的方式，计算每一位1的个数
            number += (k >> i) & 1
        }
        //最终将抵消后剩余的1放到对应的位数上
        res |= (number) % 3 << i
    }
    return res
}

//func singleNumber(nums []int) int {
//    ans := int32(0)
//    for i := 0; i < 32; i++ {
//        total := int32(0)
//        for _, num := range nums {
//            total += int32(num) >> i & 1
//        }
//        if total%3 > 0 {
//            ans |= 1 << i
//        }
//    }
//    return int(ans)
//}
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/single-number-ii/solutions/746993/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetc-23t6/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
    nums := []int{21}
    fmt.Println(singleNumber3(nums))
}

/*
https://go101.org/article/operators.html
位运算符号
11002 & 10102 results in 10002
11002 | 10102 results in 11102
11002 ^ 10102 results in 01102
11002 &^ 10102 results in 01002
*/
