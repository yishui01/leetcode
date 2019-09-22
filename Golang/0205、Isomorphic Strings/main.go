func isIsomorphic(s string, t string) bool {
    lensA := len(s)
    lensB := len(t)
    
    if lensA != lensB {
        return false
    }
    
    //s到t建立一个map映射，同时要记录t的值是否被映射了，已被映射就不能被映射为其他字符
    maps := make(map[byte]byte)
    valueMpas := make(map[byte]bool)
    
    for i:=0;i < lensA;i++ {
        if val,ok:=maps[s[i]];ok{
            if val != t[i] {
                return false
            }
        } else {
            //这个值是不是已经建立映射了
            if _,ok:=valueMpas[t[i]];ok{
                return false
            }
            
            maps[s[i]] = t[i]
            valueMpas[t[i]] = true
        }
           
    }
    
    return true
}
