package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

/*
10. Vowel
solo从小就对英文字母非常感兴趣，尤其是元音字母(a,e,i,o,u,A,E,I,O,U)，他在写日记的时候都会把元音字母写成大写的，辅音字母则都写成小写，虽然别人看起来很别扭，但是solo却非常熟练。你试试把一个句子翻译成solo写日记的习惯吧。

解答要求
时间限制：1000ms, 内存限制：100MB
输入
输入一个字符串S(长度不超过100，只包含大小写的英文字母和空格)。

输出
按照solo写日记的习惯输出翻译后的字符串S。
*/
func main() {
    inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')
    fmt.Println(vowel(input))
}
func vowel(s string) string {
    lowerS := strings.ToLower(s)
    byteS := []byte(lowerS)
    for i := range byteS {
        if byteS[i] == 'a' || byteS[i] == 'e' || byteS[i] == 'i' || byteS[i] == 'o' || byteS[i] == 'u' {
            byteS[i] = byteS[i] + 'A' - 'a'
        }
    }
    return string(byteS)
}
