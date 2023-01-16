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
删除指定目录 倒序轮询
strings.Split|-
把|-全变为实际路径
*/
/*
   // 待实现函数，在此函数中填入答题代码
   private static int delDirectorys(String delDirNameBunch, String[] dirTreeLines) {
       List<String> delDirDictionary = Arrays.asList(delDirNameBunch.split(" "));
       if (dirTreeLines.length == 1 && delDirDictionary.contains(dirTreeLines[0])) {
           return 1;
       }
       List<String> dirTreeLines2List = Arrays.asList(dirTreeLines);
       Collections.reverse(dirTreeLines2List);
       List<List<String>> childList = new ArrayList<>();
       return getDelCnt(delDirDictionary, dirTreeLines2List, childList);
   }
   private static int getDelCnt(List<String> delDirDictionary, List<String> dirTreeLines2List,
           List<List<String>> childList) {
       int delCnt = 0;
       for (String tempReverseDirLine : dirTreeLines2List) {
           String[] tempReverseDirLine2Arr = tempReverseDirLine.split("\\|-");
           int tempDirLen = tempReverseDirLine2Arr.length;
           String tempDirName = tempReverseDirLine2Arr[tempDirLen - 1];
           if (delDirDictionary.contains(tempDirName) && childList.size() == 0) {
               // 如果是删除节点，并且没有子节点(即自己是叶子节点)
               delCnt++;
           } else if (delDirDictionary.contains(tempDirName)
                   && childList.get(childList.size() - 1).size() <= tempDirLen) {
               // 如果是删除节点，并且站在当前节点角度，其没有就近且深度更深的不可删除子节点
               delCnt++;
           } else {
               // 无法删除的Dir节点加入到备选子结点列表中
               List<String> tempDirList = Arrays.asList(tempReverseDirLine2Arr);
               childList.add(tempDirList);
           }
       }
       return delCnt;
   }
*/
func delDirectorys(delDirNameBunch string, dirTreeLines []string) int {
    //倒序遍历 记下此时下标
    // 请在此补充代码
    target := strings.Split(delDirNameBunch, " ")
    for i := range dirTreeLines {
        tmp := strings.Split(dirTreeLines[i], "|-")

    }
    //倒叙
    return 0

}

func main() {
    reader := bufio.NewReader(os.Stdin)
    delDirNameBunch := readStringLine(reader)
    line := readStringLine(reader)
    num, err := strconv.Atoi(line)
    if err == nil {
        dirTreeLines := make([]string, num)
        for i := 0; i < num; i++ {
            dirTreeLines[i] = readStringLine(reader)
        }
        ans := delDirectorys(delDirNameBunch, dirTreeLines)
        fmt.Println(ans)
    }
}

func readStringLine(reader *bufio.Reader) string {
    lineBuf, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        return ""
    }
    str := strings.TrimRight(lineBuf, "\r\n")
    return str
}
