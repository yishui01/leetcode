func getMoneyAmount(n int) int {
    cache := make([][]int,n+1)
    for k,_ := range cache {
        cache[k] = make([]int,n+1)
    }
    return helper(1,n,cache)
}

func helper(start,end int, cache [][]int) int {
    if start >= end {
        return 0
    }
   if cache[start][end]!= 0 {
       return cache[start][end]
   }

    minRes := math.MaxInt32
    for i:=start;i<=end;i++{
        res := i+max(helper(i+1,end,cache),helper(start,i-1,cache))
        minRes = min(res,minRes)
    }
   cache[start][end] = minRes
    return minRes
}


func max(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func min(a,b int)int{
    if a < b {
        return a
    }
    return b
}
