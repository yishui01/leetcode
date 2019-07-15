func isInterleave(s1 string, s2 string, s3 string) bool {
    //反正不管怎么样，先递归走一遍，就当是练习递归了，很好，Time Limit Exceeded
    //return recursion(s1,s2,s3, "")
    //递归过不了之后看了答案，dp，大概理解了，把状态转移方程抄过来，注意下面的s1[i-1]这个i要减去1的原因是为了和索引对上，指的是当前字符，而不是指代上一个字符
    //dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i - 1 + j]) || (dp[i][j - 1] && s2[j - 1] == s3[j - 1 + i]);
    //这个dp还是强，递推的精髓就在于只管从前一个状态推出当前状态，不用管整体
    len1 := len(s1)
    len2 := len(s2)
    len3 := len(s3)
    
    if len1 + len2 != len3 {
        return false
    }
    
    if len1 == 0 {
        return s2 == s3
    }
    
    if len2 == 0{
        return s1 == s3
    }
    
     //row对应s1   col对应s2
    dp := make([][]bool, len2+1)
    
    for i:=0; i < len2+1; i++ {
        dp[i] = make([]bool, len1+1)
    }
   
    //先初始化最上边和最左边
    dp[0][0] = true
    for i:=1; i < len1+1; i++ { //最上边
        if dp[0][i-1] && s1[i-1] == s3[i-1] {
            dp[0][i] = true
        }
    }
    
    for i:=1; i < len2+1; i++ { //最左
        if dp[i-1][0] && s2[i-1] == s3[i-1] {
            dp[i][0] = true
        }
    }
    
    
    //开始递推
    for i := 1; i < len2+1; i++ {
        for j := 1; j <len1+1; j++ {
            if dp[i-1][j] && s2[i-1] == s3[i+j-1] || dp[i][j-1] && s1[j-1] == s3[i+j-1] {
                dp[i][j] = true
            }
        } 
    }
    
    return dp[len2][len1]
    
}



//递归大法好啊，就是过不了啊
func recursion(s1 string, s2 string, s3 string, tmpRes string) bool {
    fmt.Println(tmpRes)
    if len(s1) == 0 && len(s2) == 0 && tmpRes == s3 {
        return true
    }
    ans := false
    
    if len(s1) > 0 {
        //这里递归是在s1字符串的每个字符之间都空出一格
        ans = recursion(s1[1:], s2, s3, tmpRes + string(s1[0]))
        if ans {
            return true
        }
    }
    if len(s2) > 0 {
        //然后这里就是在s1的那每一格里面尝试所有的组合
        ans = recursion(s1, s2[1:], s3, tmpRes + string(s2[0]))
        if ans {
            return true
        }
    }
    
    return ans
    
}
