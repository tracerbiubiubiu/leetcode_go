package main

import "fmt"

/*
https://leetcode.cn/problems/reformat-the-string/
*/
//func reformat(s string) string {
//    tmpByte := make([]byte, 0)
//    length := len(s)
//    byteNu := make([]byte, 0)
//    byteAlpha := make([]byte, 0)
//    //end指向字母 start指向数字
//    for i := 0; i < length; i++ {
//        if s[i] >= '0' && s[i] <= '9' {
//            byteNu = append(byteNu, s[i])
//        } else {
//            byteAlpha = append(byteAlpha, s[i])
//        }
//    }
//    if abs(len(byteAlpha), len(byteNu)) > 1 {
//        return ""
//    }
//    i := 0
//    if len(byteNu) > len(byteAlpha) {
//        tmpByte = doAppend(tmpByte, byteNu, byteAlpha, len(byteAlpha), i)
//        tmpByte = append(tmpByte, byteNu[len(byteNu)-1])
//    } else if len(byteNu) < len(byteAlpha) {
//        tmpByte = doAppend(tmpByte, byteAlpha, byteNu, len(byteNu), i)
//        tmpByte = append(tmpByte, byteAlpha[len(byteAlpha)-1])
//    } else {
//        tmpByte = doAppend(tmpByte, byteNu, byteAlpha, len(byteNu), i)
//    }
//
//    return string(tmpByte)
//}
//func doAppend(tmpByte, byteNu, byteAlpha []byte, count, i int) []byte {
//    for i < count {
//        tmpByte = append(tmpByte, byteNu[i])
//        tmpByte = append(tmpByte, byteAlpha[i])
//        i++
//    }
//    return tmpByte
//}
func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func main() {
    s := "covid2019"
    //s := "ab123"
    fmt.Println(reformat(s))
}

func reformat(s string) string {
    length := len(s)
    tmpByte := []byte(s)
    countNu := 0
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            countNu++
        }
    }
    countAlpha := length - countNu
    if abs(countAlpha, countNu) > 1 {
        return ""
    }
    for i, j := 0, 1; i < length; i += 2 {
        if countAlpha > countNu {
            if tmpByte[i] >= 'a' && tmpByte[i] <= 'z' {
                continue
            } else {
                for tmpByte[j] >= '0' && tmpByte[j] <= '9' {
                    j += 2
                }
            }
            tmpByte[i], tmpByte[j] = tmpByte[j], tmpByte[i]
        } else {
            if tmpByte[i] >= '0' && tmpByte[i] <= '9' {
                continue
            } else {
                for tmpByte[j] >= 'a' && tmpByte[j] <= 'z' {
                    j += 2
                }
            }
            tmpByte[i], tmpByte[j] = tmpByte[j], tmpByte[i]
        }
    }
    return string(tmpByte)
}
