package main

/*
https://leetcode.cn/problems/count-integers-with-even-digit-sum/
*/
//func countEven(num int) int {
//
//    tmp := 0
//    for i:=1;i<=num;i++{
//        tmpI :=i
//        k1 := tmpI /1000
//        tmpI -= tmpI /1000*1000
//        k2 := tmpI /100
//        tmpI -= tmpI /100*100
//        k3 := tmpI /10
//        tmpI -= tmpI /10*10
//        k4 := tmpI
//        if (k1+k2+k3+k4) % 2==0{
//            tmp++
//        }
//    }
//    return tmp
//
//}
//func countEven(num int) (ans int) {
//    for i := 1; i <= num; i++ {
//        s := 0
//        for x := i; x > 0; x /= 10 {
//            s += x % 10
//        }
//        if s%2 == 0 {
//            ans++
//        }
//    }
//    return
//}


/*
小于等于num且各位数字之和为偶数
1.判断num奇偶
如果为偶，满足条件1的奇数为num/2 偶数数量为num/2
如果为奇数，奇数数量num/2+1，偶数数量num/2
2.条件2
https://leetcode.cn/problems/count-integers-with-even-digit-sum/solutions/1278781/pythongo-shu-xue-by-himymben-smnf/?languageTags=golang
*/
//func countEven(num int) int {
//    s := 0
//    for i := num / 10; i > 0; i /= 10 {
//        s += i % 10
//    }
//    ans := num / 10 * 5 - 1
//    if s % 2 == 0 {
//        ans += num % 10 / 2 + 1
//    } else {
//        ans += (num % 10 + 1) / 2
//    }
//    return ans
//}

