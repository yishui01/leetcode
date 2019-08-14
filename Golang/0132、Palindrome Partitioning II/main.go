func minCut(s string) int {
    lens:= len(s)
    if lens <=1 {
        return 0
    }
    //dp为当前字符结尾的最少分割数
    dp := make([]int, lens)
    for i:=0; i < lens; i++ {
        dp[i] = i
    }
    for i:=0; i < lens; i++ {
        findLongPal(s,i,i,dp) //奇数
        findLongPal(s,i,i+1,dp) //偶数
    }
    return dp[lens - 1];
}

//寻找s中最长的回文串
func findLongPal(s string, start,end int, dp []int) {
    for start >=0 && end < len(s) && s[start] == s[end] {
        if start == 0 { //如果已经到头了，那代表以end结尾的字符串整个就是个回文串，切割数为0
            dp[end] = 0
        } else { //没到头的话有两种情况，一种是原来的dp[end],dp[end]又有两种情况，(一种是最初始的i,还有一种是之前遍历时辐射到的end,)还有一种是当前回文串开始处的前一个字符为切割点，所以是dp[start-1]+1
            dp[end] = min(dp[end], dp[start - 1] + 1);
        }
        start--
        end++
    }
}

func min(a,b int) int {
    if a >= b {
        return b
    }
    return a
}
