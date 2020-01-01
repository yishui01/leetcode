func firstUniqChar(s string) int {

    var ints [26]int

    for _,v := range s {
        ints[v-'a']++
    }
    for k,v := range s{
        if ints[v-'a'] == 1 {
            return k
        }
    }

    return -1
   /* maps := make(map[byte]int)
    for i:=0;i<len(s);i++{
        maps[s[i]]++
    }
    for i:=0;i<len(s);i++{
        if val := maps[s[i]];val == 1 {
            return i
        }
    }

    return -1*/
}
