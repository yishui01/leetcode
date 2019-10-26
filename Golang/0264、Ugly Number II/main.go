func nthUglyNumber(n int) int {
    //dp
    dp := make([]int,n)
    dp[0] = 1
    
    l1,l2,l3 := 0,0,0
    count := 1
    
    for count < n {
        r1,r2,r3 := dp[l1]*2,dp[l2]*3,dp[l3]*5
        min := Min(r1,r2,r3)
        if r1 == min {
            l1++
        }
        if r2 == min {
            l2++
        }
        if r3 == min {
            l3++
        }
        dp[count] = min
        count++
    }
    
    return dp[count-1]
    
}

func Min(a,b,c int) int {
    if a <=b && a <= c{
        return a
    }
    if b <= a && b <= c {
        return b
    }
    if c <= a && c <= b {
        return c 
    }
    return a
}
