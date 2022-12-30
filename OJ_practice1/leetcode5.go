package main

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/
//func longestPalindrome(s string) string {
//    if len(s) <= 1 {
//        return s
//    }
//    if len(s) == 2 {
//        if s[0] == s[1] {
//            return s
//        }
//    }
//    start, end := 0, 0
//    for i := 0; i <= len(s)-1; i++ {
//        //jishu
//        left1, right1 := doExpand(s, i, i)
//        //oushu
//        left2, right2 := doExpand(s, i, i+1)
//        if right1-left1 > end-start {
//            start, end = left1, right1
//        }
//        if right2-left2 > end-start {
//            start, end = left2, right2
//        }
//
//    }
//    return s[start : end+1]
//
//}
//func doExpand(s string, left, right int) (int, int) {
//    for left > 0 && right < len(s)-1 && s[left] == s[right] {
//        left--
//        right++
//    }
//    return left, right
//}
func longestPalindrome(s string) string {
    start, end := 0, 0
    start1, end1 := 0, 0
    length := len(s)
    for i := 0; i < length; i++ {
        l1, r1 := expand(i, i, s)
        l2, r2 := expand(i, i+1, s)
        if r1-l1 > r2-l2 {
            start1, end1 = l1, r1
        } else {
            start1, end1 = l2, r2
        }
        if end1-start1 > end-start {
            start, end = start1, end1
        }
    }
    return s[start : end+1]
}
func expand(i, j int, s string) (int, int) {
    for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
        i--
        j++
    }
    return i + 1, j - 1
}
