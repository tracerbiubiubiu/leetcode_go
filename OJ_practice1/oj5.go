package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

/*

 */
//进制转换
//>1时
//十进制小数转换成二进制小数采用"乘2取整，顺序排列"法，其它进制方法类似。
//二转10 小数位前为2**n 小数位后为2**-n
func convertToN(s []string) string {
    res := "0."
    m, _ := strconv.ParseFloat(s[0], 64)
    n, _ := strconv.ParseFloat(s[1], 64)

    for count := 10; count > 0; count-- {
        tmp := m * n
        s := fmt.Sprintf("%f", tmp)
        ss := strings.Split(s, ".")
        res += ss[0]
        nu, _ := strconv.ParseFloat(ss[0], 64)
        m = tmp - nu
    }

    return res
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    question := []string{}
    for {
        input, _ := reader.ReadString('\n')
        str := strings.Fields(input)
        if str[0] == "0" && str[1] == "0" || len(str) != 2 {
            break
        }
        question = append(question, input)

    }
    for _, v := range question {
        str := strings.Fields(v)
        fmt.Println(convertToN(str))
    }
}
