package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
服务器组的均匀拆分
*/

// 待实现函数，在此函数中填入答题代码
func splitEqualTreeAndOutputNewRoot(tree []int) int {

    return -1
}

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    _, err := readIntSlice(inputReader, " ")
    if err != nil {
        return
    }

    nodes, err := readIntSlice(inputReader, " ")
    if err != nil {
        return
    }

    res := splitEqualTreeAndOutputNewRoot(nodes)
    fmt.Print(res)
}

func readIntSlice(reader *bufio.Reader, sep string) ([]int, error) {
    lineBuf, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        return nil, fmt.Errorf(err.Error())
    }

    lineBuf = strings.TrimRight(lineBuf, "\r\n")
    line := strings.Split(lineBuf, sep)
    var result []int
    for _, v := range line {
        i, err := strconv.ParseInt(v, 10, 32)
        if err != nil {
            return nil, fmt.Errorf(err.Error())
        }
        result = append(result, int(i))
    }
    return result, nil
}
