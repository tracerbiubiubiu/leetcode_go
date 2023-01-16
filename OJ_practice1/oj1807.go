package main

/*
矩阵转置
*/
import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strings"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    inputStr, err := readInput(inputReader)
    if err != nil {
        return
    }
    result := multiTranspose(inputStr)
    fmt.Println(result)
}

// 待实现函数，在此函数中填入答题代码
func multiTranspose(input string) string {
    n, ok := isValid(input)
    if !ok {
        return "ERROR"
    }
    res := ""
    for i := 0; i < n; i++ {
        for j := i; j < len(input); j += n {
            res += string(input[j])
        }
    }

    return res
}

func readInput(reader *bufio.Reader) (string, error) {
    lineBuf, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        return "", err
    }
    lineBuf = strings.TrimRight(lineBuf, "\r\n")
    lineBuf = strings.TrimRight(lineBuf, " ")
    return lineBuf, nil
}
func isValid(s string) (int, bool) {
    length := len(s)
    n := math.Sqrt(float64(length))
    if int(n)*int(n) != length {
        return 0, false
    }
    return int(n), true
}
