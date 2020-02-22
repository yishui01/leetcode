func frequencySort(s string) string {
    type node struct {
        C byte
        Val int
    }

    maps := make(map[byte]*node,len(s))
    for i:=0;i<len(s);i++{
        if r,ok := maps[s[i]];ok{
            r.Val++
        } else {
            maps[s[i]] = &node{
                C:s[i],
                Val:1,
            }
        }
    }
    
    ints := make([][]*node,len(s))
    for _,v := range maps {
        ints[v.Val-1] = append(ints[v.Val-1],v)
    }
    buf := bytes.Buffer{}
    for i:=len(ints)-1;i>=0;i--{ 
        if ints[i]!=nil{
            for k,v := range ints[i] {
                for j:=v.Val-1;j>=0;j--{
                    buf.WriteByte(ints[i][k].C)
                }
            }
        }
    }

    return buf.String()

}
