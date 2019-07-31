func numDistinct(s string, t string) int {
    //这题只会递归，然后超时，看答案，，构造二维DP数组
    //dp[i][j]有两种情况
    //1、当前字符 不相等，那么dp[i][j]还和之前一样，之前有多少子序列当前还是多少子序列 dp[i][j] = dp[i][j-1]
    //2、当前字符相等 dp[i][j] = dp[i][j-1] + dp[i-1][j-1]   //这个想了好久
    //就是说首先要把之前能构造的加上对吧，然后由于当前字符也比对成功了，那总的把这个条件用上吧，怎么用上呢，
    //就是找到前一个字符的比对情况（s的前一个字符所包含的t的前一个字符的子序列个数），找出这些序列个数有什么用？
    //这些序列个数是不是还差一个t的字符，那是不是刚好可以把目前的这个s的这个字符加进来，那不就是前边的子序列和目前的s字符
    //刚好构成当前比对的子序列所以要加上dp[i-1][j-1]
    
	if s == t {
		return 1
	}
    lenS, lenT := len(s), len(t)
    if lenT > lenS {
		return 0
	}
    
    dp := make([][]int, lenT+1)
    for i:=0; i < lenT+1; i++{
        dp[i] = make([]int, lenS+1)
    }
    for i:=0; i < lenS+1; i++ {
        dp[0][i] = 1
    }
    for i:=1; i < lenT+1; i++ {
        for j :=1; j < lenS+1; j++ {
            if t[i-1] == s[j-1] {
                dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
            } else {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[lenT][lenS]
    
}


/**

//递归回溯   Time Limit Exceeded

func numDistinct(s string, t string) int {
    var res int
    lens := len(t)
    recursion([]byte(s),[]byte(t),0,lens, &res)
    return res
}

func recursion(s, t []byte, level int, lens int, res *int) {
    if level == lens {
        *res = *res+1
        return 
    }
    
    //t[level] 为当前要找的字符
    for k,v := range s {
        if v == t[level] {fmt.Println("来了")
            recursion(s[k+1:], t, level + 1, lens, res)
        }
    }
}
*/


/**递归回溯，加上记忆数组，这种记忆数组好像只能用于返回值式的递归，将结果存在指针中的递归不能用记忆数组好像
func numDistinct(s string, t string) int {
    
    lenS := len(s)
    lenT := len(t)
    if lenS < lenT {
        return 0
    }
    memory := make([][]int, lenS)
    for i:=0; i < lenS; i++ {
        memory[i] = make([]int, lenT)
        for j:=0; j < lenT; j++ {
            memory[i][j] = -1
        }
    }
    return recursion(s,t,0,0,lenS,lenT, memory)
}

func recursion(s, t string, startS int, level int, lenS, lenT int, memory [][]int) int {
    
    if level >=lenT {
        return 1
    }
    if startS >= lenS {
        return 0
    }
    
    if memory[startS][level] != -1{
        return memory[startS][level]
    }
    var sum int
    //t[level] 为当前要找的字符
    for i:=startS; i < lenS; i++ {
        
        if s[i] == t[level] {
            sum += recursion(s, t, i+1, level + 1, lenS,lenT, memory)
        }
    }
     memory[startS][level] = sum
    return  sum
}
**/
