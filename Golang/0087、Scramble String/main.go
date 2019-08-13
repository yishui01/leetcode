func isScramble(s1 string, s2 string) bool {
    //此题有点蛋疼，看了半天不知道什么意思，自己递归写出来有问题，遂see answer，分治，将原问题不断分解为子问题
    //s1不断的切割成两段，isScramble(s1左,s2左)（s2左就是在s2中长度一样的半段）&& isScramble(s1右, s2右)  || isScramble(s1左,s2右) && isScramble(s1右边,s2左)
    lens1 := len(s1)
    lens2 := len(s2)
    
    if lens1 != lens2 || lens1 == 0 {
        return false
    }
    
    if s1 == s2 {
        return true
    }
    
    hash := make([]int, 128)
    for i:= 0; i < lens1; i++ {
        hash[s1[i]]+= 1
        hash[s2[i]]-=1
    }
    for i:= range hash {
        if hash[i] != 0 {
            return false
        }
    }
    
    for i:=1; i < lens1; i++ {
        if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]){
           return true
        }
        
        if isScramble(s1[:i], s2[lens2-i:]) && isScramble(s1[i:], s2[:lens2-i]){
            return true
        }
    }
    
    return false
}
