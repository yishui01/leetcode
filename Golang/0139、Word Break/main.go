func wordBreak(s string, wordDict []string) bool {
    //解法三：接下来用DP，设dp[i]为0-i这一段的子串是否能够被拆分
    //思路：dp[i] = dp[j] && s[j:i]为合法的单词  (j >=0 && j < i)
    //解释：每一个子串都可以被拆分成两个子串，dp[i]可以被分为两个子串，分隔点为j
    //由于j<i, 所以dp[j]已经在之前已已经求出来了，接下来只要看s[j:i]是否为合法单词就行了
    //是的话就将本索引对应的值设置为true，否则不用管，默认为false
    
    lens := len(s)
    if lens == 0 {
        return true
    }
    dp := make([]bool, lens+1) 
    dp[0] = true
    maps := make(map[string]bool)
    for _, v := range wordDict {
        maps[v] = true
    }
    for i:=0; i < lens+1; i++ {
        for j:=0; j < i; j++ {
            if dp[j]{
                if _,ok := maps[s[j:i]];ok{
                   dp[i] = true 
                }
            }
        }
    }
    return dp[lens]
}


/**
//解法二：记忆递归-----Runtime: 0 ms, faster than 100.00% 
                  -----Memory Usage: 2.2 MB, less than 100.00%  

func wordBreak(s string, wordDict []string) bool {
    maps := make(map[int]bool)
    return helper(s,wordDict, 0, maps)
}

func helper(s string, wordDict []string, start int, maps map[int]bool) bool {
       //既然递归超时，那就用记忆数组
      if s == ""{
        return true
    }
    
    //下面用到了个for range循环,每次循环代表一次新一轮的递归，那么我们实际上只要跑完第一轮的递归，几乎大部分结果都出来了
    //进行第二轮就直接拿来用就行了，这里关键点就是用什么样的索引来保存递归值，这里采用start+len(v)的方式，也就是s字符串当前的索引+合法单词长度
    //相当于以这个为起点，接下来是行还是不行（true或false），想下。比如wordDict里面有个长度为1的单词叫a，他在最外层第一轮循环时产生递归，
    //目标单词是aaaaaaaaaaaaa，那他在递归的过程中就会记录每一个数组值，因为是递归，所以应该是自底向上的，start = 13是为true, start=12是为true
    start为11时还是true，一直到start为0，最后start就是true，然后直接return，不用进行第二轮循环，这是最理想的情况，但是有可能这一层为false了，但是
    //依然能记下许多节点，供下一轮循环使用。比如这一轮记录了start为5的情况是true还是false，那么下一个单词如果长度刚好为5时，就不用递归了，直接return了
    
    
    //总结：这种循环切割流递归一般都可以用记忆数组
    
     if _, ok := maps[start]; ok {
         return maps[start]
    }
     maps[start] = false
    //一个单词一个单词的找，截取s的头部单词长度
    for _,v := range wordDict {
        if len(s) < len(v) {
            continue
        }
        if v == s[:len(v)] {
            if helper(s[len(v):],wordDict,start+len(v), maps) {
                maps[start] = true
                return true
            }
        }
    }
    
    return maps[start]
}

**/


/**
//解法一：耿直递归-----Time Limit Exceeded

func wordBreak(s string, wordDict []string) bool {
    //先递归试试,很好，TLE，至少解法是对的
    if s == ""{
        return true
    }
    
    //一个单词一个单词的找，截取s的头部单词长度
    for _,v := range wordDict {
        if len(s) < len(v) {
            continue
        }
        if v == s[:len(v)] {
            if wordBreak(s[len(v):],wordDict) {
                return true
            }
        }
    }
   return false
}
***/
