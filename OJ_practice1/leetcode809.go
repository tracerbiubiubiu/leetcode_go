package main

import "fmt"

/*
https://leetcode.cn/problems/expressive-words/
*/
func expressiveWords(s string, words []string) int {
    res := 0
    for _, v := range words {
        if canExpand(s, v) {
            res++
        }
    }
    return res
}
func canExpand(s, word string) bool {
    startS, startWord := 0, 0
    lenS, lenWord := len(s), len(word)
    for startS < lenS && startWord < lenWord {
        if s[startS] != word[startWord] {
            return false
        }
        tmp := s[startS]
        countS := 0
        for startS < lenS && s[startS] == tmp {
            startS++
            countS++
        }
        countWord := 0
        for startWord < lenWord && word[startWord] == tmp {
            startWord++
            countWord++
        }
        if countS < countWord || countS > countWord && countS < 3 {
            return false
        }
    }
    return startS == lenS && startWord == lenWord
}
func main() {
    s := "heeellooo"
    word := "hello"
    fmt.Println(canExpand(s, word))
}
