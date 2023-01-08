package week319
/*
2470
 */

func subarrayLCM(nums []int, k int) int {
    ans := 0
    n := len(nums)
    for i := range nums {
        tot := 1
        for j := i; j < n; j++ {
            tot = lcm(tot, nums[j])
            if tot == k {
                ans++
            }
            //// 剪枝：lcm 必须是 k 的因子
            if tot > k {
                break
            }
        }
    }
    return ans
}


func gcd(a,b int)int{
    if b == 0{
        return a
    }
    return gcd(b,a%b)
}
/*
最大公约数（greatest common divisor）

欧几里得辗转相除法：

gcd(x,y)表示x和y的最大公约数

进入运算时:x!=0,y!=0，本质上就是不断转换成求等价更小数的最大公约数。如果x%y=0，返回y，即最大公约数。
gcd(x,y)=gcd(y,x%y)

证明：
 设k=x/y，b=x%y  则：x=ky+b
如果n能够同时整除x和y，则(y%n)=0,(ky+b)%n=0，则b%n=0，即n也同时能够整除y和b。
由上得出：同时能够整除y和(b=x%y)的数，也必然能够同时整除x和y。故而gcd(x,y)=gcd(y,x%y)。当(b=x%y)=0，即y可以整除x，这时的y也就是所求的最大公约数了。

附上两条重要性质：gcd(a,b)=gcd(b,a)，gcd(-a,b)=gcd(a,b)


最小公倍数（least common multiple）
公式解法：
最小公倍数=两数之积/最大公约数
 */
func lcm(a,b int)int{
    return a*b/gcd(a,b)
}