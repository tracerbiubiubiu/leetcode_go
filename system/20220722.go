package main

import (
    "fmt"
    "strconv"
    "strings"
)

/*
字符编码
*/

/*
golang的内置进制转换接口
    binaryNu := strconv.FormatInt(int64(21326), 2) format是转成几进制 Parse之类的是从几进制转成10进制
    a1 := fmt.Sprintf("%b", 21326)

*/
func xutfEncoding(unicodeVal int) string {
    res := ""
    if unicodeVal < 1<<7 {
        tmp := (1 << 7) | unicodeVal
        return fmt.Sprintf("%X", tmp)
        //strconv.Formatint(int64(tmp), 16)
    }
    head := 1 << 6
    for unicodeVal >= head {
        tmp := ((1 << 6) - 1) & unicodeVal //取出后6位
        tmp = (1 << 6) | tmp
        res = strconv.FormatInt(int64(tmp), 16) + res //刚好是一字节 不用统一加完在换算16进制
        unicodeVal >>= 6
        head >>= 1 //每次末尾减6的同时，头部head右移1位
    }
    resNum := unicodeVal | head
    res = strconv.FormatInt(int64(resNum), 16) + res
    if len(res)%2 != 0 { //长度奇数补0
        res = "0" + res
    }
    return strings.ToUpper(res)
}
func main() {
    fmt.Println(xutfEncoding(21326))
    fmt.Println(xutfEncoding(34))
    fmt.Println(xutfEncoding(1225859))
}
