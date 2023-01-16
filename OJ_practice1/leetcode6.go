package main

import "fmt"

/*
https://leetcode.cn/problems/zigzag-conversion/description/
*/
//优化的话可以不用二维数组
func convert(s string, numRows int) string {
    tmp := make([][]byte, numRows)
    length := numRows
    if length == 1 {
        return s
    }
    count := len(s)
    for i := 0; i < length; i++ {
        tmp = append(tmp, []byte{})
    }
    group := 2*numRows - 2
    for i := 0; i < count; i++ {
        //cur := i
        carry := min(i%group, group-i%group)
        //carry := cur % numRows
        tmp[carry] = append(tmp[carry], s[i])
    }
    res := ""
    for i := range tmp {
        tmpS := string(tmp[i])
        //fmt.Println(tmpS + "\n")
        res += tmpS
    }
    return res
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
func main() {
    s := "PAYPALISHIRING"
    numRows := 4
    fmt.Println(convert(s, numRows))
}
