package week319
/*
2472
 */

func maxPalindromes(s string, k int) int {
    res :=0
    for i := range s{
        res+=findH(s,i,i,k)
        res+=findH(s,i,i+1,k)
    }
    return res
}
func findH(s string,i,j,k int)int{
    res := 0
    if i==j{
        tmp := i
        for i >0{
            if s[tmp-1]==s[tmp+1]{
                i++
            }
        }
        if (i-tmp)*2+1>=k{
            res++
        }
    }else {
        left,right := i,j
        if s[left]!=s[right]{
            return res
        }
        for left>=0&&right<=len(s)-1{
            if s[left-1]==s[right+1]{
                left--
                right++
            }
        }
        if right-left+1>=k{
            res++
        }
    }
    return res
}