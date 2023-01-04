package main

import (
    "math/big"
)

/*

 */
//a数值很大 最大为2的10**5-1 要使用大数包
func addBinary(a string, b string) string {

    ai, _ := new(big.Int).SetString(a, 2)
    bi, _ := new(big.Int).SetString(b, 2)

    ai.Add(ai, bi)
    return ai.Text(2)
}

