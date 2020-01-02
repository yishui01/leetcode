func findTheDifference(s string, t string) byte {

    ssum,tsum := 0,0
    for _,x := range s {
        ssum += int(x)
    }
    for _,x := range t {
        tsum += int(x)
    }
    return byte(tsum-ssum)

   /* maps := make(map[byte]int)

    for i:=0;i<len(s);i++{
        maps[s[i]]++
    }

    for i:=0;i<len(t);i++{
        if _,ok := maps[t[i]];ok{
            maps[t[i]]--
        } else {
            return t[i]
        }
    }

    for k,v := range maps {
        if v != 0 {
            return k
        }
    }

    return ' '*/
}
