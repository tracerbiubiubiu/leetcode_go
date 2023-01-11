package main

import (
    "fmt"
    "strconv"
)

/*
外观数列
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。



1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//func countAndSay(n int) string {
//    var res string
//    if n == 1 {
//        return "1"
//    }
//    for i := 2; i <= n; i++ {
//        res = countAndSayHelper(countAndSay(i - 1))
//    }
//    return res
//}
//
////由n-1生成n
////或门 一个为true就是true
//func countAndSayHelper(str string) string {
//  ans := strings.Builder{}
//  n := len(str)
//  j := 0
//  for i := 0; i <= n; i++ {
//      //fmt.Println("***")
//      //fmt.Println(i)
//      if i == n || str[j] != str[i] {
//          count := i - j
//          ans.WriteString(strconv.Itoa(count))
//          ans.WriteByte(str[j])
//          j = i
//      }
//  }
//  //fmt.Println(reflect.TypeOf(ans))
//  return ans.String()
//}

func main() {
    a := countAndSay(6)
    fmt.Println(a)
}

func countAndSay(n int) string {
    var res string
    if n == 1 {
        res = "1"
    } else {
        res = handle(countAndSay(n - 1))
    }
    return res
}

//由n-1得出n
func handle(s string) string {
    var partn []byte
    length := len(s)
    //Subpart := []
    count := 1
    if length == 1 {
        return "11"
    }
    for i := 1; i < length; i++ {
        if s[i] != s[i-1] {
            part := strconv.Itoa(count)+string(s[i-1])
                part1 := []byte(part)
            for i := range part1 {
                partn = append(partn, part1[i])
            count = 1
            }
        } else {
            count += 1
        }
        if i == length-1 {
            part := strconv.Itoa(count)+string(s[i])
            part1 := []byte(part)
            for i := range part1 {
                //fmt.Println(i)
                partn = append(partn, part1[i])}
        }
    }
    return string(partn)
}
