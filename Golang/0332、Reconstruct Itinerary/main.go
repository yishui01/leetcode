func findItinerary(tickets [][]string) []string {
    maps := make(map[string][]string)
    for _,v := range tickets {
        maps[v[0]] = append(maps[v[0]],v[1])
    }
    for k,_ := range maps{
        sort.Strings(maps[k])
    }

    res := make([]string,0)

    dfs(maps,"JFK",&res)

    //因为是dfs讲元素添加进res的，所有实际上是从最后一个元素开始添加的，因此需要反转
    for i:=0;i<len(res)/2;i++{
        res[i],res[len(res)-1-i] = res[len(res)-1-i],res[i]
    }

    return res
    
}

func dfs(maps map[string][]string,from string,res *[]string) {
    for len(maps[from]) > 0 {
        nextForm := maps[from][0]
        maps[from] = maps[from][1:]
        dfs(maps,nextForm,res)
    }
    *res = append(*res,from)
}

