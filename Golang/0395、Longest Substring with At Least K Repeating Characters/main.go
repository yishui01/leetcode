func longestSubstring(s string, k int) int {
    return helper(s,k)
}

//看的答案，分而治之，一个字符串中如果存在某个字符x并且x出现的次数小于k，那么最长的子串必定不包含x
//由此可以将该字符串分割成三段   0....x-1 , x , x+1...n
//对左右两段分别求解
func helper(s string,k int) int {
    if s == "" || len(s) < k{
        return 0
    }
    var maps [26]int 
    for i:=0;i<len(s);i++{
        maps[s[i]-'a']++
    }

    //方法一：执行用时：52ms
    /*for i:=0;i<len(s);i++{
        if maps[s[i]-'a'] < k {
            //这里是分割点
            //s[:i]和s[i+1:]
            leftAns := helper(s[:i],k)
            if len(s[i+1:]) < leftAns {
                return leftAns
            }
            return max(leftAns,helper(s[i+1:],k))
        }
    }
    return len(s)*/

    //方法二：执行用时：0ms
    //这个是一段一段求的，少了字符串截取成两段时，右边那段在底层复制的过程，方法一他不停的打断，底层其实在不停的复制，那么右半段一直在被重复的复制，方法二就不会
    ok := true
    start := 0
    res := 0;
    for i:=0;i<len(s);i++{
        if maps[s[i]-'a'] < k {
           res = max(res, longestSubstring(s[start:i], k));
           ok = false;
           start = i + 1;
        }
    }
    if ok {
        return len(s)
    }
    return  max(res, longestSubstring(s[start:len(s)], k))
}

func max(a,b int)int{
    if a > b {
        return a
    }
    return b
}

