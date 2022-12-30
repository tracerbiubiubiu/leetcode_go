
/*
让我们来玩个字符消除游戏吧, 给定一个只包含大写字母的字符串s，消除过程是如下进行的：

1)如果s包含长度为2的由相同字母组成的子串，那么这些子串会被消除，余下的子串拼成新的字符串。
例如”ABCCBCCCAA”中”CC”,”CC”和”AA”会被同时消除，余下”AB”, “C”和”B”拼成新的字符串”ABBC”。

2)上述消除会反复一轮一轮进行，直到新的字符串不包含相邻的相同字符为止。
例如”ABCCBCCCAA”经过一轮消除得到”ABBC”，再经过一轮消除得到”AC”

输入
第一行输入一个正整数 T(1 <= T <= 50) , 表示有 T 组测试数据.

对于每个测试数据输入包括一行, 由大写字母组成的字符串s，长度不超过100.

输出
对于每组测试数据, 若最后可以把整个字符串全部消除, 就输出 Yes, 否则输出 No.

输入样例 1 复制

2
ABCCBA
ABCCCCCBBBBB
输出样例 1

Yes
No
提示
多重暴力搜索就行, 遇到两个连在一起的 就删掉 循环直到再也没有两个连在一起为止.
*/

package main

import "fmt"
func Purified(s string)string{
	Length :=len(s)
	if Length==1{return s}
	var b string
	var flag int=1

	for flag==1{
		b=s
		if len(b)<=1{flag=0}//缺少判断会死循环
		for i:=0;i<len(b)-1;i++{//这块取值范围 <len(b)-1  就是取到len(b)-2 +1 刚好 我真是个傻逼
			if s[i]==s[i+1]{
				s=b[:i]+b[i+2:]
				break
			}
			if i==len(b)-2{
				flag=0
				break
			}
		}
	}
	return b
}
//func Purified(s string)string{
//	Length :=len(s)
//	if Length==1{return s}
//	//var newString string=""
//	LOOP:
//		for i:=0;i<len(s)-2;i++{
//			if len(s)<=1{return s}
//			if s[i]==s[i+1]{
//				s=s[:i]+s[i+2:]
//				goto LOOP
//			}
//		}
//		return s
//}
func main(){
	var a int
	var s string
	c:=make([]string,0)
	fmt.Scanln(&a)
	for i:=0;i<a;i++{
		fmt.Scanln(&s)
		c=append(c,s)
	}

	for i,_:=range c{
		if len(Purified(c[i]))>0{
			fmt.Println("no")
		}else {fmt.Println("yes")}
	}
}

