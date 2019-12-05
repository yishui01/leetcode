func integerBreak(n int) int {
    maps := make(map[int]int)
    return dfs(n,maps)
}

//用dfs试试
func dfs(num int,maps map[int]int) int {
    if num == 1 {
        return 1
    }
    if val,ok := maps[num];ok {
        return val
    }
    res := 0
    for i:=1;i<num;i++ {
        res = max(res,max(i*(num-i),i * dfs(num-i,maps)))
    }
    maps[num] = res
    return res
}

func max(a, b int)int{
    if a > b {
        return a
    }
    return b
}
