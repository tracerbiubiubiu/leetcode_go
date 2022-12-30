package main

/*
给你一个 正 整数 num ，输出它的补数。补数是对该数的二进制表示取反。

给定的整数 num 保证在 32 位带符号整数的范围内。
num >= 1
你可以假定二进制数不包含前导零位。

异或
0101
^
0111
=
0010

只需要构造一个比num大的全1mask序列即可

func findComplement(num int) int {
    i := 1
    for i <= num {
        i <<= 1
    }
    return i - num - 1
}


*/

/*思路很简单
输入为 5 是 101
输出为 2 是 010
上下两个加起来就是 111 即2的3次方-1

输入为 8 是 1000
输出为 7 是 0111
上下两个加起来就是 1111 即2的4次方-1

同理
就可以知道
找到一个比num大的 2的n次幂的数 本题为a
然后结果就是a - num - 1

作者：interesting-jangwfa
链接：https://leetcode-cn.com/problems/number-complement/solution/shu-zi-de-bu-shu-qiao-jie-by-interesting-kdum/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
//func main()  {
//    n := 6
//    a := convertToBin(n)
//    fmt.Println(a)
//}

//位运算
func findComplement(num int) int {
    i := 1
    for i <= num {
        i <<= 1
    }
    return (i - 1) ^ num
}

/*
func findComplement(num int) int {
    binNum := convertToBin(num)

    s := ""
    for _,v := range binNum{
        if v == 0 {
            s += "1"
        }else {
            s += "0"
        }
    }
    intNum,_ := strconv.Atoi(s)
    return intNum
}
func convertToBin(num int) string {
    s := ""

    if num == 0 {
        return "0"
    }

    // num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
    for ; num > 0; num /= 2 {
        lsb := num % 2
        // strconv.Itoa() 将数字强制性转化为字符串
        s = strconv.Itoa(lsb) + s
    }
    return s
}
*/