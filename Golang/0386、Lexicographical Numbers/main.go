func lexicalOrder(n int) []int {
    res := make([]int,0)
    for  i:=1;i<=9;i++ {
        dfs(i,n,&res);
    }
    return res
}

func dfs(cur,n int,res *[]int) {
    if cur > n {
        return 
    }
    *res = append(*res,cur)
    cur *= 10
    for i:=0;i<10;i++{
        dfs(cur+i,n,res)
    }
}

