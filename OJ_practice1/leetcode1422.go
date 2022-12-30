package main

import "strings"

/*
一个由若干0和1组成的字符串inputStr，分割成左右两个非空子字符串后，定义字符串的【权重】为左子字符串中0的数量加上右自字符串中1的数量
请计算并返回字符串分割后的最大权重

输入：inputStr = "100010"
输出：4
解释：
将字符串 inputStr 划分为两个非空子字符串的可行方案有：
左子字符串 = "1" ，右子字符串 = "00010"，权重 = 0 + 1 = 1，
左子字符串 = "10" ，右子字符串 = "0010"，权重 = 1 + 1 = 2，
左子字符串 = "100" ，右子字符串 = "010"，权重 = 2 + 1 = 3，
左子字符串 = "1000" ，右子字符串 = "10"，权重 = 3 + 1 = 4，
左子字符串 = "10001" ，右子字符串 = "0"，权重 = 3 + 0 = 3。
所以最大权重为 4 。

输入：inputStr = "111111"
输出：5
*/
func maxWeight(inputStr string) int {
    score := strings.Count(inputStr[:1], "0") + strings.Count(inputStr[1:], "1")
    max := score
    for i := 1; i <= len(inputStr)-1; i++ {
        if inputStr[i] == '1' {
            score--
        } else {
            score++
        }
        max = maxOf2(max, score)
    }
    return max
}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}
