package main

func minPersonRequired(monitorDatas [][]int, threshold []int) int {
    tmp := []int{}
    for i := range monitorDatas {
        count := 0
        for j := range threshold {
            if monitorDatas[i][j] > threshold[j] {
                count++
            }
        }
        tmp = append(tmp, count)
    }
    num := map[int]int
    tmp2 := []int{}
    for i := range tmp {
        if tmp[i] > 1 {
            tmp2 = append(tmp2, 5)
        } else if tmp[i] == 1 {
            tmp2 = append(tmp2, 2)
        }
    }
    res := len(tmp2)
    for i := range tmp2 {
        if i+tmp2[i] <= len(tmp2)-1 {
            if _,ok := num[tmp[i]];!ok{
                num[tmp[i]] = i+tmp[i]
            }
            for i,v := range num{
                if num[i]==v{
                    res--
                    delete(num,tmp[i])
                }

            }


            continue
        }

    }

    return res
}
