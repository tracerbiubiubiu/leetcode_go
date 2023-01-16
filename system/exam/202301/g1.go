package main

import (
    "fmt"
    "strings"
    "unicode"
)

//func wordStat(inputs []string) int {
//    //把多行先拼接为1行，在统计数量最后是否要添加换行
//    tmp := strings.Join(inputs, "\n")
//    //tmp := strings.ReplaceAll(tmp, "\\", "")
//    count := 0
//    flag := false
//    i := 0
//    for ; i < len(tmp); i++ {
//        for i < len(tmp) && unicode.IsLetter(rune(tmp[i])) {
//            flag = true
//            i++
//        }
//        for i < len(tmp) && ((tmp[i-1] == '\\' && tmp[i] == '\n') || (tmp[i] == '\\' && tmp[i-1] == tmp[i])) {
//            i++
//        }
//        if flag {
//            count++
//            flag = false
//        }
//    }
//
//    return count
//}
func main() {
    inputs := []string{

        "a \\",
        "aagg\\",
        "\\",
        "aa,b.",
        "h\\",
        "ell\\",
        "o\\",
        //"engineers , worldwide, ,.courses     part",
        //"y.. .vendors.",
        //",hell\\",
        //"o",
    }
    fmt.Println(wordStat(inputs))
}

//
func wordStat(inputs []string) int {
    res := 0
    //tmp := strings.Join(inputs, "")
    //fmt.Printf("%v\n", tmp)
    tmpStr := ""
    newInput := []string{}
    for i := range inputs {
        if strings.HasSuffix(inputs[i], "\\") {
            tmpStr += inputs[i][:len(inputs[i])-1]
            if i == len(inputs)-1 {
                newInput = append(newInput, tmpStr)
            }
        } else {
            newInput = append(newInput, tmpStr+inputs[i])
            tmpStr = ""
        }
    }
    fmt.Println(newInput)
    flag := false
    //fmt.Println(newInput[2])
    for i := range newInput {
        for j := 0; j < len(newInput[i]); {
            for j < len(newInput[i]) && !unicode.IsLower(rune(newInput[i][j])) {
                j++
            }
            //now := j
            for j < len(newInput[i]) && unicode.IsLower(rune(newInput[i][j])) {
                flag = true
                j++
            }
            if flag {
                res++
                flag = false
            }
        }
    }
    return res
}
