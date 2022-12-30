package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
题目描述
在我们的windows系统中，有一个命令行程序叫做命令提示符（也即cmd.exe）。
现在小华正在学习cmd的使用方法，可是小华用的并不是windows的系统，所以他很困扰。
他现在学习了cd命令的使用方法，希望你能够帮助小华编写一个程序，模拟cd命令的响应，让小华能够练习使用cd命令。

cd命令的格式与效果：
cd a //进入a文件夹
cd .. //返回上一级文件夹，进入到某一个盘符的根目录时，cd ..命令将不会再改变当前路径

假设cmd初始的目录为：“C:\Users\hp”

以下为命令的样例
首先程序输出“C:\Users\hp>”
小华输入命令“cd ..”
程序响应，输出“C:\Users>”
小华输入命令“cd hp”
程序响应，输出“C:\Users\hp>”
小华输入命令“cd Desktop”
程序响应，输出“C:\Users\hp\Desktop>”

解答要求
时间限制：1000ms, 内存限制：100MB
输入
第一行一个整数n（n<=1000），表示小华将要输入的命令数。
接下来n行，每行一个命令，保证cd与参数之间有且只有一个空格。

输出
模拟出cmd程序应有的响应。

样例
输入样例 1 复制

7
cd ..
cd hp
cd Desktop
cd ..
cd ..
cd ..
cd ..
输出样例 1

C:\Users\hp>
C:\Users>
C:\Users\hp>
C:\Users\hp\Desktop>
C:\Users\hp>
C:\Users>
C:\>
C:\>
提示样例 1
*/
//还是字符串操作 map
func CmdLine(s string)string{
	//var b []string
	b:=strings.Split(s," ")
	fmt.Println(b[1])
	var curLoc string
	curLoc="C:\\Users\\hp>"
	switch b[1] {
	case "hp":
		curLoc="C:\\Users\\hp>"
		fmt.Println(curLoc)
	case "Desktop":
		curLoc="C:\\Users\\hp\\Desktop>"
		fmt.Println(curLoc)
	case "..":
		switch curLoc {
		case  "C:\\Users\\hp>":
			curLoc="C:\\Users>"
			fmt.Println(curLoc)
		case "C:\\Users\\hp\\Desktop>":
			curLoc="C:\\Users\\hp>"
			fmt.Println(curLoc)
		case "C:\\Users>":
			curLoc="C:\\>"
			fmt.Println(curLoc)
		default:
			curLoc="C:\\>"
			fmt.Println(curLoc)
		}
	default:
		}
	return curLoc
}
func main(){
	var a int
	x:=make([]string,0)
	fmt.Println("C:\\Users\\hp>")
	fmt.Scanln(&a)
	//var s string
	for i:=1;i<=a;i++{
		s:=bufio.NewReader(os.Stdin)
		input,err:=s.ReadString('\n')
		//x=append(x,*s)
	}
	for i,_:=range x{
		CmdLine(x[i])
	}

}