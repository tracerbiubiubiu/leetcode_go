package main

/*
2021-3-26 可信考试科目一工作级 golang
http://3ms.huawei.com/km/blogs/details/10082051
*/

/*
一、第一题描述
给定一个表示时间的字符创timeStr，格式为hh:mm(hh为小时，mm为分钟)，其中某几位数字被隐藏（用？表示）。
其有效时间为00:00到23:59之间的所有时间，包括00:00和23:59,00:00最早， 23:59最晚
请替换timeStr中隐藏的数字，得到最晚的有效时间，以格式hh:mm输出
示例
输入：timeStr = "1?:0?"
输出："19:09"
解释：以 '1' 开头的最晚一小时是 19 ，以 '0' 开头的最晚一分钟是 09 。

输入：timeStr = "??:59"
输出："23:59"
提示： ?的个数范围是[0:4]
https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits/

*/
func maximumTime(time string) string {
    tmpStr := []byte(time)
    if tmpStr[0] == '?' {
        if tmpStr[1] > '3' && tmpStr[1] <= '9' {
            tmpStr[0] = '1'
        } else {
            tmpStr[0] = '2'
        }
    }
    if tmpStr[1] == '?' {
        if tmpStr[0] <= '1' {
            tmpStr[1] = '9'
        } else {
            tmpStr[1] = '3'
        }
    }
    if tmpStr[3] == '?' {
        tmpStr[3] = '5'
    }
    if tmpStr[4] == '?' {
        tmpStr[4] = '9'
    }
    return string(tmpStr)
}
