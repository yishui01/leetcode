func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    res := 0
    i,j := 0,0 

    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {
            i++
            res++
        }
        j++
    }

    return res

    /*
    cindex := 0
    for i:=0;i<len(g);i++{
        for j:=cindex;j<len(s);j++{
            cindex++
            if s[j] >= g[i] {
                res++
                break
            }
        }
    }

    return res*/
}
