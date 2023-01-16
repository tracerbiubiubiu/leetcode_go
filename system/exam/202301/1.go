/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 * Description: 上机编程认证
 * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
 * 只能import Go标准库
 */
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

type department struct {
	city           string
	departmentName string
	personNum      int64
}

func sortForRemove(departments []department) {
	sort.Slice(departments, func(i, j int) bool {
		if departments[i].city == departments[j].city {
			if departments[i].personNum == departments[j].personNum {
				return departments[i].departmentName < departments[j].departmentName
			}
			return departments[i].personNum > departments[j].personNum
		}
		return departments[i].city < departments[j].city
	})
}

func removeLessFive(departments []department) []department {
	ans := make([]department, 0)
	ans = append(ans, departments[0])
	lastCity := departments[0].city
	maxNum := 1
	for _, dep := range departments[1:] {
		if dep.city == lastCity {
			maxNum += 1
		} else {
			maxNum = 1
			lastCity = dep.city
		}
		if maxNum < 6 {
			ans = append(ans, dep)
		}
	}
	return ans
}

func sortForResult(departments []department) {
	sort.Slice(departments, func(i, j int) bool {
		if departments[i].city == departments[j].city {
			return departments[i].departmentName < departments[j].departmentName
		}
		return departments[i].city < departments[j].city
	})
}

func getTop5Department(departments []department) []department {
	sortForRemove(departments)
	ans := removeLessFive(departments)
	sortForResult(ans)
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
