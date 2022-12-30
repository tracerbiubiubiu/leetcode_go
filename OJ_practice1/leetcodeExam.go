package main

import (
    "fmt"
)

/*
假设有搅乱顺序的一群儿童成一个队列。 每个儿童由一个整数对表示，
其中 w 是这个儿童 的体重，k 是排在这个儿童前面且体重大于或等于 w 的儿童数量。 编写一个算法来重建这个队 列。
输入:
[[8,0], [4,4], [8,1], [5,0], [6,1], [5,2]]

输出:
解释:[5,2]前面两个儿童的体重分别是 5 和 8，且只有两个儿童;[6,1]前面只有[8,0]儿童的体重 大于他/她，并且不能和[5,2]换位置，否则会导致[5,2]的 2 不对。
[[5,0], [8,0], [5,2], [6,1], [4,4], [8,1]]
*/
//func reconstructQueue(people [][]int) [][]int {
//
//}

/*
请你按照指定的规则对一个非空字符串 中的字符进行重新组合，并返回重新组合后的结果。
给定一个整数，使得重新组合后的字符串中相同字母的位置间隔距离interval
如果无法完成重新组合，则返回一个空字符串 ""。
*/
/*
input = "xxyyzz", interval = 3
"xyzxyz"
*/
//试试双指针
func combineChars(input string, interval int) string {
    length := len(input)
    if length < interval || length == 0 {
        return ""
    }
    tmp := []byte(input)
    tmpMap := make(map[byte]int)
    //先分组
    for i := 0; i < length; i++ {
        if _, ok := tmpMap[tmp[i]]; ok {
            tmpMap[tmp[i]]++
        } else {
            tmpMap[tmp[i]] = 1
        }
    }
    return string(res)
}
func main() {
    a := "aaazxxyy"
    //  a := "xxyyzz"
    res := combineChars(a, 2)
    fmt.Println(res)
}
