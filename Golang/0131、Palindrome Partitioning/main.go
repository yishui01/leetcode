func partition(s string) [][]string {
    res := make([][]string, 0)
    helper(s, []string{}, &res)
    return res
}

func helper( s string ,  tmpRes []string,   res *[][]string) {
    lens := len(s)
    if lens == 0 {
        dist := make([]string, len(tmpRes))
        copy(dist, tmpRes)
        *res = append(*res, dist)
        return 
    }
    
    tmpLens := len(tmpRes)
    for i:=0; i < lens; i++ {
        if checkPalin(s[:i+1]) {
            //当前是回文
            tmpRes = append(tmpRes, s[:i+1])
            helper(s[i+1:],tmpRes, res)
            tmpRes = tmpRes[:tmpLens]
        }
    }
    
}

func checkPalin(s string) bool {
    lens := len(s)
    for i:=0; i < lens/2; i++ {
        if s[i] != s[lens-1-i] {
            return false
        }
    }
    
    return true
}
