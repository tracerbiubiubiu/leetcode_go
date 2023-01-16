package main

import "fmt"

/*

 */
func printTable(n, m int) {

    tmpRow := "---+"
    row := "+"
    for i := 0; i < m; i++ {
        row += tmpRow
    }
    tmpCol := ""
    for i := 0; i < 4*m+1; i++ {
        if i%4 == 0 {
            tmpCol += "|"
        } else {
            tmpCol += " "
        }
    }

    col := fmt.Sprintf("%s", tmpCol)
    fmt.Println(len(col))
    for i := 0; i < 2*n+1; i++ {
        if i%2 == 0 {
            fmt.Println(row)
        } else {
            fmt.Println(col)
        }
    }
}
func main() {
    var m int
    var n int
    fmt.Scanf("%d %d", &n, &m)
    printTable(n, m)
}
