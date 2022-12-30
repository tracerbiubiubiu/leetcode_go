package main

import "fmt"

/*
https://leetcode.cn/problems/integer-to-roman/
*/
//for range无序 应该变成有序
type roman struct {
    val int
    roa string
}

func intToRoman(num int) string {
    res := ""
    var romanMap []roman = []roman{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    for i := range romanMap {
        for num != 0 {
            for num-romanMap[i].val >= 0 {
                res += romanMap[i].roa
                num -= romanMap[i].val
            }
            break

        }
    }
    return res
}
func main() {
    a := 999
    fmt.Println(intToRoman(a))
}
