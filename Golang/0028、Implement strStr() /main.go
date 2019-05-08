func strStr(haystack string, needle string) int {
    subLens := len(needle)
    lens := len(haystack)
    res := -1
    
    if subLens == 0 {
        return 0
    }
    
    if lens == 0 || lens < subLens{
        return -1
    }
    times := lens - subLens + 1
    for i:=0; i < times; i++ {
        if needle[0] == haystack[i] {
            //开始匹配
            isFind := true
            z := i
            for j:=0; j < subLens; j++ {
                if needle[j] != haystack[z] {
                    isFind = false
                    break
                }
                z++
            }
            if isFind {
                return i
            }
        }
    }
    
    return res
    
}
