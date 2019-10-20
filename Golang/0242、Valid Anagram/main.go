func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    maps := make(map[byte]int)
    
    for i:=0;i<len(s);i++{
        maps[s[i]]++
        maps[t[i]]--
    }
    
    for _,v := range maps {
        if v != 0 {
            return false
        }
    }
    
    return true
    
}

// func isAnagram(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }
    
//     maps := make(map[byte]int)
   
//     for i:=0;i<len(s);i++{
//        maps[s[i]]++
//     }
    
//     for i:=0;i<len(t);i++{
//         if val,ok := maps[t[i]];ok && val > 0 {
//             maps[t[i]]--
//         } else {
//             return false
//         }
//     }
    
//     return true
// }
