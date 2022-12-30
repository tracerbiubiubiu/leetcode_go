package main

import (
    "fmt"
)

/*
https://leetcode.cn/problems/string-to-integer-atoi/
*/
//题目很恶心 没描述清楚 看别人的用例才完全明白
func myAtoi(s string) int {
    length := len(s)
    sign := 1
    res := 0
    i := 0
    for ; i < length; i++ {
        if s[i] == ' ' {
            continue
        }
        break
    }
    if i < length {
        if s[i] == '-' {
            sign = -1
            i++
        } else if s[i] == '+' {
            i++
        }
    }

    for i < length && s[i] >= '0' && s[i] <= '9' {
        res = res*10 + int(s[i]-'0')
        if sign*res > 1<<31-1 {
            return 1<<31 - 1
        } else if sign*res < -1<<31 {
            return -1 << 31
        }

        i++
    }
    return sign * res

}
func main() {
    a := "21474836460"
    fmt.Println(myAtoi(a))
}

//func myAtoi(s string) int {
//    abs, sign, i, n := 0, 1, 0, len(s)
//    //丢弃无用的前导空格
//    for i < n && s[i] == ' ' {
//        i++
//    }
//    //标记正负号
//    if i < n {
//        if s[i] == '-' {
//            sign = -1
//            i++
//        } else if s[i] == '+' {
//            sign = 1
//            i++
//        }
//    }
//    for i < n && s[i] >= '0' && s[i] <= '9' {
//        abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
//        if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
//            return math.MinInt32
//        } else if sign*abs > math.MaxInt32 {
//            return math.MaxInt32
//        }
//        i++
//    }
//    return sign * abs
//}

//func myAtoi(s string) (ans int) {
//    state := "start"
//    table := map[string][]string{
//        "start":  []string{"start", "signed", "number", "end"},
//        "signed": []string{"end", "end", "number", "end"},
//        "number": []string{"end", "end", "number", "end"},
//        "end":    []string{"end", "end", "end", "end"},
//    }
//    sign := 1   // 符号部分
//    number := 0 // 数字部分
//    getIdx := func(c byte) int {
//        switch c {
//        case ' ':
//            return 0 // start为table[0]
//        case '-', '+':
//            return 1 // signed为table[1]
//        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
//            return 2 // number为table[2]
//        default:
//            return 3 // end为table[3]
//        }
//        return 3
//    }
//
//    for i := range s {
//        state = table[state][getIdx(s[i])] // to_state = table[from_state][new_char_idx]
//        if state == "number" {
//            number = number*10 + (int(s[i]) - '0')
//            if sign == 1 {
//                number = min(number, math.MaxInt32)
//            } else if sign == -1 { // 因为math.MaxInt32和-math.MinInt32并不相等（相差1）
//                number = min(number, -math.MinInt32)
//            }
//        } else if state == "signed" {
//            if s[i] == '-' {
//                sign = -1
//            }
//        }
//    }
//    return sign * number
//}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
