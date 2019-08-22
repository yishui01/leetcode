func wordBreak(s string, wordDict []string) []string {
    maps := make(map[string][]string)
    return helper(s,wordDict,maps)
}

//这题很难，要想清楚，递归+记忆数组，这题和前一题是经典问题，要反复理解
func helper(s string, wordDict []string, maps map[string][]string)[]string{
    if s == ""{
        return []string{""} //里面加个空元素很重要，递归起始值
    }
    if val, ok := maps[s]; ok {
        return val
    }
    res := make([]string, 0)
    
    for _, v := range wordDict{
        if len(s) >= len(v) && s[:len(v)] == v {
            rem := helper(s[len(v):], wordDict, maps)
              for _, val := range rem{
                  if val == ""{
                      res = append(res, v)
                  } else {
                     res = append(res, v+" "+val)
                  }
            }
        }
    }
    maps[s]=res
    return res
}
