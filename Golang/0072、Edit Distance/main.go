func minDistance(word1 string, word2 string) int {
    //看了大佬的解法dp, 这种只要求值的一般都是dp
    //设dp[i][j]为word1的前i和字符转换成word2前j个字符需要的次数
    //dp[i][j] = word1[i] == word2[j] ? dp[i-1][j-1] : (min(dp[i][j-1], dp[i-1][j-1], dp[i-1][j-1]) + 1)
    //思路：如果当前比较的字符相等，那很明显，dp[i][j]就等于dp[i-1][j-1]，但是如果不等呢
    //这里有三条路径
    //第一个就是dp[i-1][j] //代表i-1个字符构成j所需要的组合数，在此基础上删掉当前的i字符
    //第二个是的dp[i][j-1] //i个字符构成j-1的组合数 ，现在只构成了j-1，还差一个，添加一个j字符
    //第三个是dp[i-1][j-1] //i-1构成了j-1个字符，在当前基础上把i字符替换为j字符即可
    //上面的删除、添加、替换都是一步操作，所以要加一，最后表达式为  min(dp[i][j-1], dp[i-1][j-1], dp[i-1][j-1]) + 1
    
    //这样从i想i-1可能比较难理解，现在从i想i+1
    //dp[i][j] 现在要达到dp[i+1][j+1]
    //在不等的情况下，
    //1、替换，把当前不等的字符替换就行，dp[i][j] + 1
    //2、添加，dp[i+1][j] //用i+1构成j时（此时目标是构成j+1）。那么只需要在此基础上 加上当前的 j+1号位的字符 dp[i+1][j]+1
    //3、删除，dp[i][j+1] //用i构成了j+1，提前达到了目标，可是当前字符是i+1，多了个字符，所以要删掉多余的字符 dp[i][j+1]+1
     
    lens1 := len(word1)
    lens2 := len(word2)
    
    if lens1 == 0 {
        return lens2
    }
    
    if lens2 == 0 {
        return lens1
    }
    
    dp := make([][]int, lens1+1)
    
    //初始化dp边界
    for i:=0; i < lens1+1; i++ {
        dp[i] = make([]int, lens2+1)
        if i == 0 {
            for j:=0; j < lens2+1; j++ {
                dp[i][j] = j
            }
        } else {
            dp[i][0] = i
        }
    }
    
    for i:=1; i < lens1+1; i++ {
        for j:=1; j < lens2+1; j++ {
            if word1[i-1] == word2[j-1] { 
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i][j-1],dp[i-1][j], dp[i-1][j-1]) + 1
            }
        }
    }
    
    return dp[lens1][lens2]
    
}

func min(a,b,c int) int {
    if a <= b && a <= c {
        return a
    } else if b <= a && b <= c {
        return b
    }
    return c
}
