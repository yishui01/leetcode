func findRepeatedDnaSequences(s string) []string {
    //
    var codes ['T' + 1]uint32
    codes['A'] = 0
    codes['C'] = 1
    codes['G'] = 2
    codes['T'] = 3

    var repeated []string
    seen := make([]byte, 1<<20)
    var cur uint32
    for i := 0; i < len(s); i++ {
        cur = ((cur << 2) & 0xFFFFF) | codes[s[i]]
        if i < 9 {
            continue
        }
        if seen[cur] == 0 {
            seen[cur] = 1
        } else if seen[cur] == 1 {
            repeated = append(repeated, s[i-9:i+1])
            seen[cur] = 2
        }
    }

    return repeated
}

// func findRepeatedDnaSequences(s string) []string {
//     //看了下Rabin-Karp，感觉不适合这道题，这道题要找重复出现过的，不是直接字符串匹配
//     maps := make(map[string]bool)
//     res := make([]string,0)
//     lens := len(s)
//     for i:=0;i<lens-9;i++{
//         tmp := s[i:i+10]
//         if val,ok:=maps[tmp];ok{
//             if !val{
//                 res = append(res,tmp)
//                 maps[tmp] = true
//             }
//         } else {
//             maps[tmp] = false
//         }
//     }
//     return res
// }
