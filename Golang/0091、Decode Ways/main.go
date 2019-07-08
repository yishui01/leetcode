func numDecodings(s string) int {
    //有一个结论，dfs适合求解所有的结果集，dp适合求解所有结果数目
    //所以本题用dp来做，dfs写在下面的注释中
    //dp[i]表示当前长度为i的字符串共有多少种解法，想想当前长度的组合数和之前的长度组合数有什么关联吗
    //要到达当前长度，有几种方法,首先可以由dp[i-1]加一个字符来达到（如果加的字符有效的话），还有dp[i-2]加两个字符达到当前长度，如果满足一定条件的话
    //这不就是跳台阶吗 dp[i] = dp[i-1] + dp[i-2]
    
    lens := len(s)
    if lens == 0 {
        return 0
    }
    if s[0] == '0' {
        return 0
    }
    if lens == 1 {
        return 1
    }
    
    dp := make([]int, len(s)+1)
    dp[0] = 1
    dp[1] = 1
    
    for i:=2; i <= lens; i++ {
        if s[i-1] >= '1' && s[i-1] <= '9' { //这里s[i-1]实际上就是当前字符，因为dp的那个i代表的是长度值，长度值转换成索引要减掉1
            dp[i] += dp[i-1]
        }
        if (s[i-2] == '1') || (s[i-2] == '2' && s[i-1] >='0' && s[i-1] <='6') {
            dp[i] += dp[i-2]
        }
    }
    
    return dp[lens]
    
}


/**
func numDecodings(s string) int {
    //回溯大法好 Runtime:2092 ms 哈哈哈哈
    res :=0
    if s[0] == '0' {
        return 0
    }
    recursion(s, "", len(s), 0, &res)
    return res
}

func recursion(s string, tmpS string, lens int, tmpLens int, res *int) {
    if tmpS != "" {
        if val,ok := strconv.Atoi(tmpS);   ok == nil && val <= 26  && val > 0 && tmpS[0] != '0' {
            if lens == tmpLens {
                *res = *res+1
                return 
            }
        } else {
           return
        }
    }
    
     //这里的s为当前截取过的s
    if s[0] != '0' {
        recursion(string(s[1:]), string(s[0]), lens, tmpLens+1, res)
        if len(s) >=2 {
         recursion(string(s[2:]), string(s[:2]), lens, tmpLens+2, res)
        }
    }
    
}
**/
