package main

import "strings"

/*
三、第三题描述
请设计一个建议的垃圾邮件过滤系统，主要实现三个功能：
Add(string keyword)向系统中添加关键字， 若系统中已存在关键字keyword, 则不做任何处理，最后返回所有关键字个数；
Remove(string keyword) 删除系统中关键字keyword，然后返回生意关键字个数， 若系统中不存在关键字keyword,则不作任何处理并返回-1；
说明 Add和Remove函数中关键字keyword仅包含小写字符；
Filter(string keyword) 对语句sentence的每个单词按如下规则进行一遍关键字过滤，并最终返回过滤后的语句，如过滤后没有剩余单词，请返回空字符串""；
*语句sentense格式，由单空格分割，首位无空格的若干单词组成，单词仅包含小写字母；
*过滤是指对语句中的单词前缀进行匹配，并删除匹配到的内容，过滤后的语句也要符合sentense格式（注意：单空格分割，首位无空格）。如单词"abcabcdef",过滤一遍关键字"abc"后为"abcdef"
若语句中某单词的前缀能够匹配到系统中的对个关键字，则选择最长的关键字进行过滤。

输入：
["FilterSystem","add","filter","remove"]
[[],["hello"],["helloworld hi"],["hello"]]
输出：[null,1,"world hi",0]

解释：
FilterSystem obj = FilterSystem();
obj.Add("hello"); // 添加关键字 hello，返回 1；
obj.Filter("helloworld hi"); // 对单词 helloworld 进行一遍关键字过滤，前缀匹配到关键字 hello，hello 被过滤，过滤后为world；
对单词 hi 也进行一遍关键字过滤，未前缀匹配到任何关键字;
因此返回"world hi"；
obj.Remove("hello"); // 删除关键字 hello，返回删除后的剩余关键字个数 0；
注：输出中的 null 表示此对应函数无输出（等同于 C 语言的 void 类型）

输入：
["FilterSystem","add","add","filter","remove","add","add","add","filter","remove","filter"]
[[],["he"],["hello"],["helloworld hi"],["he"],["hi"],["abc"],["hi"],["hi hellohello hi worldhello hi"],["he"],["hi"]]
输出：[null,1,2,"world hi",1,2,3,3,"hello worldhello",-1,""]

解释：
FilterSystem obj = FilterSystem();
obj.Add("he"); // 返回 1；
obj.Add("hello"); // 返回 2；
obj.Filter("helloworld hi"); // 对单词 helloworld 进行一遍关键字过滤，前缀匹配到两个关键字 he、hello，其中 hello 是最长的关键字，因此 hello 被过滤，过滤后为world；
对单词 hi 也进行一遍关键字过滤，未前缀匹配到任何关键字;
最终返回 "world hi"；
obj.Remove("he"); // 返回 1；
obj.Add("hi"); // 返回 2；
obj.Add("abc"); // 返回 3；
obj.Add("hi"); // 返回 3；
obj.Filter("hi hellohello hi worldhello hi"); // 3个单词hi均前缀匹配到关键字hi，过滤后为""；
单词 hellohello 前缀匹配到hello，过滤后为hello；
worldhello 未前缀匹配到任何关键字，不进行过滤；
去除首尾空格，单词间仅保留一个空格，最终返回 "hello worldhello"；
obj.Remove("he"); // 返回 -1；
obj.Filter("hi"); // 返回空字符串 ""。
注：输出中的 null 表示此对应函数无输出（等同于 C 语言的 void 类型）
*/
type FilterSystem struct {
    dataCenter map[string]bool
}

func (filterSystem *FilterSystem) filter(sentence string) string {
    ans := ""
    strArr := strings.Split(sentence, " ")
    for _, str := range strArr {
        maxPri := ""
        for key, _ := range filterSystem.dataCenter {
            if strings.HasPrefix(str, key) && len(key) > len(maxPri) {
                maxPri = key
            }
        }
        if strings.TrimPrefix(str, maxPri) != "" {
            ans += strings.TrimPrefix(str, maxPri) + " "
        }
    }
    return strings.Trim(ans, " ")
}
