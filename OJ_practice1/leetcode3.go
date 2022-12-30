package main

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
func lengthOfLongestSubstring(s string) int {
    
}