package main

import (
    "fmt"
    "strconv"
    "strings"
)

/*
http://3ms.huawei.com/km/groups/3803117/blogs/details/13430679?l=zh-cn#top
*/

//func displayPages(pageCount, maxWidth, currentPage int) string {
//    res := make([]string, maxWidth)
//    if pageCount <= maxWidth {
//        for i := 1; i <= pageCount; i++ {
//            res = append(res, strconv.Itoa(i))
//        }
//        return strings.Join(res, "")
//    }
//    //首尾页码按钮无区别
//    //761 1...4567 1234...7 有区别么
//    if currentPage == 1 || currentPage == maxWidth {
//        for i := 1; i <= maxWidth-2; i++ {
//            res = append(res, strconv.Itoa(i))
//        }
//        tail := "..." + strconv.Itoa(pageCount)
//        res = append(res, tail)
//        return strings.Join(res, "")
//    }
//
//    mid := currentPage
//    head, tail := mid, mid
//    blankLen := maxWidth - 3
//    flag := 0
//    isHead, isTail := false, false
//    for blankLen >= 3 {
//        if head-1 <= 1 {
//            tail++
//            blankLen--
//            continue
//        }
//        if pageCount-tail <= 1 {
//            head--
//            blankLen--
//            continue
//        }
//        if flag == 0 {
//            head--
//            flag = 1
//        } else {
//            tail++
//            flag = 0
//        }
//
//        blankLen -= 1
//    }
//    if head-1 <= 1 {
//        isTail = true
//    }
//    if pageCount-tail <= 1 {
//        isHead = true
//    }
//    if isHead {
//        headInfo := "1" + "..."
//        res = append(res, headInfo)
//        for i := head; i <= maxWidth; i++ {
//            res = append(res, strconv.Itoa(i))
//        }
//        res = append(res, strconv.Itoa(pageCount))
//        return strings.Join(res, "")
//    } else if isTail {
//        tailInfo := "..." + strconv.Itoa(pageCount)
//        for i := 1; i <= tail; i++ {
//            res = append(res, strconv.Itoa(i))
//        }
//        res = append(res, tailInfo)
//        return strings.Join(res, "")
//    } else if !isTail && !isHead {
//        tmp := ""
//        headInfo := "1" + "..."
//        tailInfo := "..." + strconv.Itoa(pageCount)
//        for i := head; i <= tail; i++ {
//            tmp += strconv.Itoa(i)
//        }
//        res = append(res, headInfo, tmp, tailInfo)
//        return strings.Join(res, "")
//    }
//    return ""
//}
//还有一种先总结...出现在左边与右边的点，在进行判断，共判断4次，为特殊情况全部打印，左边有，右边有，两边都有
//判断的点为x>p-m/2与x<m/2
func displayPages(pageCount, maxWidth, currentPage int) string {
    res := make([]string, maxWidth)
    if pageCount <= maxWidth {
        for i := 1; i <= pageCount; i++ {
            res = append(res, strconv.Itoa(i))
        }
        return strings.Join(res, "")
    }
    head, tail := currentPage, currentPage
    flag := 0
    for tail-head+1 < maxWidth-4 {
        if flag == 0 {
            head--
            flag = 1
            continue
        } else {
            tail++
            flag = 0
            continue
        }
    }
    tmp := make([]string, pageCount)
    for i := head; i <= tail; i++ {
        tmp = append(tmp, strconv.Itoa(i))
    }
    if head != 1 {
        tmp[1] = "..."
    }
    if tail != maxWidth {
        tmp[tail-2] = "..."
    }
    tmp = tmp[1 : tail-1]
    return "1" + strings.Join(tmp, "") + strconv.Itoa(maxWidth)
}
func main() {
    a := displayPages(7, 6, 1)
    fmt.Println(a)
}

/*
func displayPages(pageCount, maxWidth, currentPage int) string {

    res := make([]string, 0, maxWidth)

    if pageCount <= maxWidth {

        for i := 1; i <= pageCount; i++ {

            res = append(res, strconv.Itoa(i))

        }

        return strings.Join(res, " ")

    }

    var left, right = currentPage, currentPage

    var flag = true

    for right-left+1 < maxWidth {

        if flag {

            flag = false

            if left > 1 {

                left--

            }

        } else {

            flag = true

            if right < pageCount {

                right++

            }

        }

    }

    for i := left; i <= right; i++ {

        res = append(res, strconv.Itoa(i))

    }

    headBuf := currentPage - left

    tailBuf := right - currentPage

    if left != 1 {

        res[0] = "1"

        if res[1] != strconv.Itoa(currentPage){

        res[1] = "..."

        headBuf--

        }

    }

    if right != pageCount {

        res[maxWidth-1] = strconv.Itoa(pageCount)

        if res[maxWidth-2] != strconv.Itoa(currentPage){

        res[maxWidth-2] = "..."

        tailBuf--

        }

    }

    if headBuf > tailBuf {

        res[1] = "..."

    }

    if tailBuf > headBuf {

        res[maxWidth-2] = "..."

    }

    return strings.Join(res, " ")

}
*/
