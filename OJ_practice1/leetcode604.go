package main

import (
    "fmt"
    "strconv"
    "strings"
)

/*
https://leetcode.cn/problems/solve-the-equation/
*/
func solveEquation(equation string) string {
    _, res := canDoCalculate(equation)
    return res
}
func canDoCalculate(s string) (bool, string) {
    tmpS := strings.Replace(s, "-", "+-", -1)
    tmp := strings.Split(tmpS, "=")
    left := strings.Split(tmp[0], "+")
    right := strings.Split(tmp[1], "+")
    leftX, leftNu := doRange(left)
    rightX, rightNu := doRange(right)
    calcLeft := leftX - rightX
    calcRight := rightNu - leftNu
    if calcLeft == 0 {
        if calcRight == 0 {
            return false, "Infinite solutions"
        }
        return false, "No solution"
    }

    if calcRight%calcLeft != 0 {
        return false, "No solution"
    }

    return true, "x=" + strconv.Itoa(calcRight/calcLeft)

}
func doRange(s []string) (int, int) {
    numX, numNu := 0, 0
    for _, v := range s {
        if strings.Contains(v, "x") {
            x := strings.Split(v, "x")[0]
            if x == "" || x == "+" {
                x = "1"
            }
            if x == "-" {
                x = "-1"
            }
            nuX, _ := strconv.Atoi(x)
            numX += nuX
        } else {
            nu, _ := strconv.Atoi(v)
            numNu += nu
        }
    }
    return numX, numNu
}
func main() {
    //a := "x+5-3+x=6+x-2"
    a := "-x=-1"
    b := solveEquation(a)
    fmt.Println(b)
}
