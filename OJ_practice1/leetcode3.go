package main

import "fmt"

/*


 */
//func lengthOfLongestSubstring(s string) int {
//    res := 0
//    length := len(s)
//    if length <= 1 {
//        return length
//    }
//    for i := 0; i < length; i++ {
//
//        j := i + 1
//        tmpMap := make(map[byte]int)
//        tmpMap[s[i]] = i
//        for j < length {
//            if _, ok := tmpMap[s[j]]; ok {
//
//                res = maxOf2(res, len(tmpMap))
//                break
//            } else {
//                tmpMap[s[j]] = j
//                j++
//            }
//
//        }
//        res = maxOf2(res, len(tmpMap))
//
//    }
//    return res
//}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}
//另一种思路，直接更新map的值指向的下标，start到下标之间的key全部删掉
func lengthOfLongestSubstring(s string) int {
    res := 0
    tmpMap := make(map[byte]int)
    length := len(s)
    for start,end :=0,0; end < length;end++{
        if _,ok := tmpMap[s[end]];ok{
            for start < tmpMap[s[end]]{
                delete(tmpMap,s[start])
                start++
            }
            start = tmpMap[s[end]]+1

        }

        tmpMap[s[end]] = end
        res = maxOf2(res,end-start+1)
    }
    return res
}
func main() {
    a := "abba"
    //a := "abcabcbb"
    fmt.Println(lengthOfLongestSubstring(a))
}