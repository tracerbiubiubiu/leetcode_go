package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

//bufentongguo
/*
7
shenzhen ebg 5
shenzhen dbg 7
shenzhen cbg 5
shenzhen abg 6
shenzhen zbg 5
shenzhen bbg 4
beijing cbg 1
*/
type department struct {
    city           string
    departmentName string
    personNum      int64
}

//先从每个城市选最多5个
//先把每个城市筛选出来，在逐一对比
//先选择人数最多
//人数相等按字典序
func getTop5Department(departments []department) []department {
    ans := make([]department, 0)
    allCity := make(map[string]int)

    for _, v := range departments {
        if _, ok := allCity[v.city]; !ok {
            allCity[v.city] = 1
        } else {
            allCity[v.city]++
        }
    }
    city2 := make([]string, 0)
    for i := range allCity {
        city2 = append(city2, i)
    }
    //按城市字典升序
    sort.Strings(city2)
    //每个城市建一个数组
    tmpCity := make([][]department, len(city2))
    //按照city的排序，对tmpcity进行填充
    for city := range city2 {
        for i := range departments {
            if departments[i].city == city2[city] {
                tmpCity[city] = append(tmpCity[city], departments[i])
            }
        }
    }
    //填充完后进行排序,前大后小
    for city := range tmpCity {
        sort.Slice(tmpCity[city], func(i, j int) bool {
            if tmpCity[city][i].personNum > tmpCity[city][j].personNum {
                return true
            }
            return false
        })
    }

    //循环结束，得到一个有序的二维数组
    //依次遍历数组，如果长度>5只取到第五个
    for i := range tmpCity {

        tmpAns := make([]department, 0)
        if len(tmpCity[i]) <= 5 {
            tmpAns = append(tmpAns, tmpCity[i]...)
        } else {
            tmpAns = append(tmpAns, tmpCity[i][:5]...)
        }
        //还要对这五个进行字典排序
        tmpAnsMap := make(map[string]department)
        tmpAnsMapSlice := make([]string, 0)
        for i := range tmpAns {
            tmpAnsMap[tmpAns[i].departmentName] = tmpAns[i]
            tmpAnsMapSlice = append(tmpAnsMapSlice, tmpAns[i].departmentName)
        }

        sort.Strings(tmpAnsMapSlice)
        for _, v := range tmpAnsMapSlice {
            ans = append(ans, tmpAnsMap[v])
        }

    }

    return ans

}

func main() {
    var n int
    if _, err := fmt.Scanf("%d\n", &n); err != nil {
        return
    }

    departments := make([]department, n)
    reader := bufio.NewReader(os.Stdin)
    for i := 0; i < n; i++ {
        var err error
        departments[i], err = readline(reader)
        if err != nil {
            return
        }
    }
    ans := getTop5Department(departments)
    for _, v := range ans {
        fmt.Println(v.city, v.departmentName)
    }
}

func readline(reader *bufio.Reader) (department, error) {
    lineBuf, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        return department{}, fmt.Errorf(err.Error())
    }

    lineBuf = strings.TrimRight(lineBuf, "\r\n")
    line := strings.Split(lineBuf, " ")

    num, err := strconv.ParseInt(line[2], 10, 32)
    if err != nil {
        return department{}, fmt.Errorf(err.Error())
    }

    return department{line[0], line[1], num}, nil
}
