package main

import "strings"

/*
https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
*/
func isPrefixOfWord(sentence string, searchWord string) int {
    index := strings.Index(sentence, searchWord)
    if index == -1 {
        return index
    }
    splitS := strings.Split(sentence, " ")
    for i, v := range splitS {
        if strings.HasPrefix(v, searchWord) {
            return i + 1
        }

    }
    return -1
}
