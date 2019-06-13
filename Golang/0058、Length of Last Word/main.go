func lengthOfLastWord(s string) int {
    //这题还有个巧妙的解法，他不是求最后一个单词的长度吗，那就TMD就直接从后往前遍历就行了啊，不一定要从前往后啊，我去，大佬真是666
    lens := len(s)
    if lens == 0 {
        return 0
    }
    res := 0
    flag := 0
    for i :=0; i < lens; i++ {
        if s[i] == ' ' {
            flag = 1
        } else {
            if flag == 1 {
                res = 0
            }
            res++
            flag = 0
        }
    }
    
    return res
    
}
